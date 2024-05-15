package message

import (
	"github.com/rs/zerolog"
)

type message struct {
	log       *zerolog.Event
	text      string
	flushFunc func()
}

func newMessage(event *zerolog.Event, text string, flushFunc func()) *message {
	return &message{
		log:       event,
		text:      text,
		flushFunc: flushFunc,
	}
}

func (m *message) Send() {
	m.log.Msg(m.text)

	if m.flushFunc != nil {
		m.flushFunc()
	}
}
