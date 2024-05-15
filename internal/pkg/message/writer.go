package message

import "github.com/rs/zerolog"

type writer struct {
	logger    *zerolog.Logger
	flushFunc func()
}

func newWriter(logger *zerolog.Logger, flushFunc func()) *writer {
	return &writer{
		logger:    logger,
		flushFunc: flushFunc,
	}
}
