package tzkt

import (
	"github.com/dipdup-net/go-lib/tzkt/api"
	"github.com/dipdup-net/go-lib/tzkt/events"
)

// Message -
type Message struct {
	Type  events.MessageType
	Level uint64
	Body  []api.BigMapUpdate
}

func newMessage() Message {
	return Message{
		Body: make([]api.BigMapUpdate, 0),
	}
}

func (msg *Message) clear() {
	msg.Level = 0
	msg.Type = 0
	msg.Body = make([]api.BigMapUpdate, 0)
}

func (msg Message) copy() Message {
	message := newMessage()
	message.Type = msg.Type
	message.Level = msg.Level

	for i := range msg.Body {
		message.Body = append(message.Body, msg.Body[i])
	}

	return message
}
