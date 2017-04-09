package hub

import (
	"log"
	"github.com/timojarv/gochat/message"
)

type Hub struct {
	clients ClientList
	Broadcast chan message.Message
	Register chan Client
}

type Client interface {
	Send(message.Message)
}

type ClientList map[Client]bool


// Generate a hub (central message relay)
func New() *Hub {
	return &Hub{
		clients: make(ClientList),
		Broadcast: make(chan message.Message),
		Register: make(chan Client),
	}
}

// Run the central message relay
func (hub *Hub) Run() {
	for {
		select {
			case client := <- hub.Register:
				if _, ok := hub.clients[client]; ok {
					delete(hub.clients, client)
					log.Println("Hub.Run: a client unregistered")
				} else {
					hub.clients[client] = true
					log.Println("Hub.Run: a new client registered")
				}
			case message := <- hub.Broadcast:
				for client := range hub.clients {
					go client.Send(message) // Abstract away the channel logic to make efficient bots possible
				}
		}
	}
}