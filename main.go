package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/message"
	"github.com/tbd-paas/platform-bootstrapper/internal/pkg/server"
)

// generateLogs generates log messages
func generateLogs(logger message.Logger, stop context.CancelFunc) {
	defer stop()

	// simulate a process with a beginning and end
	for i := 0; i <= 10; i++ {
		// write to the channel if we have one open otherwise just generically log
		time.Sleep(1 * time.Second)

		logger.Info(fmt.Sprintf("%d", i))
	}
}

func main() {
	// start the message processor
	ctx, stop := context.WithCancel(context.Background())

	logger := message.NewLogger(ctx, zerolog.DebugLevel)

	go logger.Start()
	logger.Info("starting logger...")

	// start the process
	go generateLogs(logger, stop)

	// start the server
	srv := server.NewServer(logger, ctx)
	srv.Start()

	// stop the logger
	logger.Shutdown()
}
