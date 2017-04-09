// Package client implements the worker that interacts with a front-end client
package client

import (
	"log"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/message"
	"github.com/timojarv/gochat/user"
	"github.com/timojarv/gochat/webtoken"
	"github.com/gorilla/websocket"
)

// Client implements the client interface (from package hub)
type Client struct {
	ws *websocket.Conn
	broadcast chan message.Message
	hub *hub.Hub
	send chan message.Message
	err chan bool
	user *user.User
}

// CreateClient is a factory returning a new Client instance bound to a specific websocket
func New(ws *websocket.Conn) *Client {
	return &Client{
		ws: ws,
		send: make(chan message.Message),
		err: make(chan bool),
	}
}

// Run starts the Client instance
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

// Send takes care of sending the broadcast messages to the socket
func (client *Client) Send(message message.Message) {
	// Only feed if user is authenticated
	if client.user != nil {
		client.send <- message // Simply feed to the outgoing channel
	}
}

// Listen listens on the websocket for incoming messages
func (client *Client) Listen() {
	// Start listening to the websocket
	var received message.Message

	for {
		// Read from websocket
		received = message.Message{}
		err := client.ws.ReadJSON(&received)
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				log.Println("Client.Listen: websocket closed by client")
				client.err <- true
				return
			} else {
				log.Println("Client.Listen: ", err)
			}
		}

		if client.user != nil {
			// Let's don't let user spoof their name
			received.Sender = client.user.Username
			client.broadcast <- received
			continue
		}

		// Handle the handshake (authentication)
		userId, err := webtoken.GetSubject(received.Body)
		if err != nil {
			client.send <- message.Error(err.Error())
			continue
		}

		// Set the user this client is dealing with
		client.user, err = user.FindById(userId)
		if err != nil {
			client.send <- message.Error(err.Error())
			continue
		}

		// Tell the user their name to confirm the authentication as being succesful
		client.send <- message.Meta(client.user.Username)
	}
}


// Register reports the Client's existence to the Hub
func (client *Client) Register(hub *hub.Hub) {
	client.hub = hub
	client.broadcast = hub.Broadcast
	hub.Register <- client
}

// Unregister reports shutdown of the Client to the Hub
func (client *Client) Unregister() {
	if client.hub == nil {
		log.Println("Client.Unregister: client is not registred")
		return
	}

	client.hub.Register <- client
	client.ws.Close()
}