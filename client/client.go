package client

import (
	"log"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/message"
	"github.com/gorilla/websocket"
)

// Implements the client interface (from package chat)
type Client struct {
	ws *websocket.Conn
	broadcast chan message.Message
	hub *hub.Hub
	send chan message.Message
	err chan bool
}

// Client generator
func CreateClient(ws *websocket.Conn) *Client {
	return &Client{
		ws: ws,
		send: make(chan message.Message),
		err: make(chan bool),
	}
}

// Create client routine for a websocket
func (client *Client) Run(hub *hub.Hub) {
	client.Register(hub)
	defer client.Unregister()

	go client.Listen() // Start the websocket listener

	// Check for errors and relay messages to the websocket
	for {
		select {
			case message := <- client.send:
				client.ws.WriteJSON(message)
			case <- client.err:
				return
		}
	}
	
}

// Listen to the message relay
func (client *Client) Send(message message.Message) {
	client.send <- message // Simply feed to the outgoing channel
}

func (client *Client) Listen() {
	// Start listening to the websocket
	var message message.Message

	for {
		// Read from websocket
		err := client.ws.ReadJSON(&message)
		if err != nil {
			log.Println("Client.Run: websocket connection closed")
			client.err <- true
			return
		}

		client.broadcast <- message
	}
}

func (client *Client) Register(hub *hub.Hub) {
	client.hub = hub
	client.broadcast = hub.Broadcast
	hub.Register <- client
}

func (client *Client) Unregister() {
	if client.hub == nil {
		log.Println("Client.Unregister: client is not registred")
		return
	}

	client.hub.Register <- client
	client.ws.Close()
}