package main

import "time"

const MessageKind = iota + 1

type Message struct {
	Kind      uint32    `json:"kind"`
	ID        string    `json:"ID"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created-at"`
}

func newMessage(id string, body string, createdAt time.Time) *Message {
	return &Message{
		Kind:      MessageKind,
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
