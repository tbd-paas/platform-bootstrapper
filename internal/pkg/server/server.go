package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
)

type Server struct {
	Context  context.Context
	Stop     context.CancelFunc
	Log      *zerolog.Logger
	Instance *http.Server
	Stream   chan string
}

func NewServer(debug bool) *Server {
	ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)

	// setup the logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if debug {
		logger.Level(zerolog.DebugLevel)
	}

	return &Server{
		Context: ctx,
		Stop:    stop,
		Log:     &logger,
		Stream:  make(chan string, 100),
		Instance: &http.Server{
			Addr: ":8080",
		},
	}
}

func (server *Server) Run() {
	defer server.Stop()

	// start the http service
	go func() {
		if err := server.Instance.ListenAndServe(); err != http.ErrServerClosed {
			server.Log.Error().Msgf("unable to start http server: %v", err)
		}
	}()

	// listen for any interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// block until we receive a signal
	for {
		select {
		case <-signalChan:
			server.Log.Info().Msg("received interrupt signal, shutting down...")

			// gracefully shutdown
			if err := server.Instance.Shutdown(server.Context); err != nil {
				server.Log.Error().Msgf("unable to gracefully shutdown http server: %v", err)
			}

			server.Log.Info().Msg("http server stopped")

			return
		case <-server.Context.Done():
			return
		}
	}
}
