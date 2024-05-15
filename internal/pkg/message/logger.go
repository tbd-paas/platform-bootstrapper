package message

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"

	"github.com/rs/zerolog"
)

type Logger interface {
	Start()
	Shutdown()

	AddWriterFor(string, io.Writer, func())
	RemoveWriterFor(string)

	Info(string)
	Error(string)
	Warn(string)
	Debug(string)
}

type logger struct {
	lock sync.Mutex

	messages chan *message
	context  context.Context
	writers  map[string]*writer
	level    zerolog.Level
}

func NewLogger(ctx context.Context, level zerolog.Level) *logger {
	messageLogger := &logger{
		lock:     sync.Mutex{},
		context:  ctx,
		messages: make(chan *message),
		level:    level,
		writers:  map[string]*writer{},
	}

	// setup the mandatory initial writer.  this ensures that we always have a writer
	// regardless.
	messageLogger.AddWriterFor("stdout", os.Stdout, nil)

	return messageLogger
}

func (l *logger) Start() {
	// listen for any interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for {
		select {
		case msg := <-l.messages:
			msg.Send()
		case <-signalChan:
			// received interrupt; stop writing log messages
			break
		case <-l.context.Done():
			// context is done; stop writing log messages
			break
		}
	}
}

func (l *logger) Shutdown() {
	l.Info("stopping logger...")
	defer close(l.messages)
}

func (l *logger) AddWriterFor(name string, writer io.Writer, flushFunc func()) {
	l.lock.Lock()
	defer l.lock.Unlock()

	logger := zerolog.New(writer).With().Timestamp().Logger().Level(l.level)

	l.writers[name] = newWriter(&logger, flushFunc)
}

func (l *logger) RemoveWriterFor(name string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	delete(l.writers, name)
}

func (l *logger) GetContext() context.Context {
	return l.context
}

func (l *logger) Info(text string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for info := range l.writers {
		l.messages <- newMessage(l.writers[info].logger.Info(), text, l.writers[info].flushFunc)
	}
}

func (l *logger) Error(text string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for err := range l.writers {
		l.messages <- newMessage(l.writers[err].logger.Error(), text, l.writers[err].flushFunc)
	}
}

func (l *logger) Warn(text string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for warn := range l.writers {
		l.messages <- newMessage(l.writers[warn].logger.Warn(), text, l.writers[warn].flushFunc)
	}
}

func (l *logger) Debug(text string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for debug := range l.writers {
		l.messages <- newMessage(l.writers[debug].logger.Debug(), text, l.writers[debug].flushFunc)
	}
}
