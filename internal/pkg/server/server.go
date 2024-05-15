package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/message"
)

const (
	interruptMessage = "received interrupt signal, shutting down..."
	finishedMessage  = "bootstrapper process finished, shutting down..."
)

type server struct {
	context  context.Context
	instance *http.Server
	logger   message.Logger
}

func NewServer(logger message.Logger, ctx context.Context) *server {
	return &server{
		context:  ctx,
		instance: &http.Server{Addr: ":8081"},
		logger:   logger,
	}
}

func (server *server) initialize() {
	// route the /logs path to the stream logs function
	http.HandleFunc("/logs", server.logs)

	go func() {
		if err := server.instance.ListenAndServe(); err != http.ErrServerClosed {
			panic(fmt.Errorf("unable to start http server: %v", err))
		}
	}()
	server.logger.Info(fmt.Sprintf("http server started on address [%s]...", server.instance.Addr))
}

func (server *server) Start() {
	server.initialize()
	defer server.Shutdown()

	// listen for any interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// block until we receive a signal
	for {
		select {
		case <-signalChan:
			server.logger.Info(interruptMessage)

			return
		case <-server.context.Done():
			server.logger.Info(finishedMessage)

			return
		}
	}
}

func (server *server) Shutdown() {
	server.logger.Info("http server stopping...")

	// gracefully shutdown
	if err := server.instance.Shutdown(server.context); err != nil {
		server.logger.Info(fmt.Sprintf("unable to gracefully shutdown http server: %v", err))
	}

	server.logger.Info("http server stopped...")
}

// logs streams logs that are written to a channel for a duration that the event stream is open.
func (server *server) logs(response http.ResponseWriter, request *http.Request) {
	// set headers for server sent events
	response.Header().Set("Content-Type", "text/event-stream")
	response.Header().Set("Cache-Control", "no-cache")
	response.Header().Set("Connection", "keep-alive")
	response.Header().Set("Access-Control-Allow-Origin", "*")

	// add an http logger to the message processor.  this allows this http process to write logs
	// that are output via the response.
	flushFunc := func() {
		flusher, ok := response.(http.Flusher)
		if ok {
			flusher.Flush()
		} else {
			server.logger.Error("flusher not supported")
		}
	}

	server.logger.AddWriterFor("http", response, flushFunc)
	defer server.logger.RemoveWriterFor("http")

	// start log streaming
	for {
		select {
		case <-server.context.Done():
			return
		case <-request.Context().Done():
			return
		}
	}
}
