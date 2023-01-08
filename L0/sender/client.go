package main

import "github.com/gorilla/websocket"

type Client struct {
	setup    *Setup
	id       int
	socket   *websocket.Conn
	outbound chan []byte
}

func newClient(setup *Setup, socket *websocket.Conn) *Client {
	return &Client{
		setup:    setup,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (client *Client) write() {
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (client Client) close() {
	client.socket.Close()
	close(client.outbound)
}
