package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
},
}

type Setup struct {
	clients    []*Client
	nextID     int
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func newSetup() *Setup {
	return &Setup{
		clients:    make([]*Client, 0),
		nextID:     0,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}
