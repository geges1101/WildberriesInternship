package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
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

func (setup *Setup) run() {
	for {
		select {
		case client := <-setup.register:
			setup.onConnect(client)
		case client := <-setup.unregister:
			setup.onDisconnect(client)
		}
	}
}

func (setup *Setup) broadcast(message interface{}, skip *Client) {
	data, _ := json.Marshal(message)
	for _, c := range setup.clients {
		if c != skip {
			c.outbound <- data
		}
	}
}

func (setup *Setup) send(message interface{}, client *Client) {
	data, _ := json.Marshal(message)
	client.outbound <- data
}

func (setup *Setup) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "upgrading error", http.StatusInternalServerError)
		return
	}
	client := newClient(setup, socket)
	setup.register <- client

	go client.write()
}

// add new client
func (setup *Setup) onConnect(client *Client) {
	log.Println("client connected: ", client.socket.RemoteAddr())

	setup.mutex.Lock()
	defer setup.mutex.Unlock()
	client.id = setup.nextID
	setup.nextID++
	setup.clients = append(setup.clients, client)
}

// delete given client
func (setup *Setup) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	client.close()
	setup.mutex.Lock()
	defer setup.mutex.Unlock()

	target := -1
	for i, c := range setup.clients {
		if c.id == client.id {
			target = i
			break
		}
	}

	copy(setup.clients[target:], setup.clients[target+1:])
	setup.clients[len(setup.clients)-1] = nil
	setup.clients = setup.clients[:len(setup.clients)-1]
}
