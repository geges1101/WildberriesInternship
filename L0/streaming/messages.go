package streaming

import "time"

type Message interface {
	Key() string
}

type NewMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (m *NewMessage) Key() string {
	return "data.created"
}
