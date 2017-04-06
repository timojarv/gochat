package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/client"
	"github.com/timojarv/gochat/bot"
)

var upgrader = websocket.Upgrader{}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Init a new message relay
	hub := hub.CreateHub()

	// Websocket route adds new client to the relay
	http.HandleFunc("/ws", addClient(hub))

	// Start the message relay
	go hub.Run()

	// Initialize the bot
	bot.CreateBot("gopher", hub)

	// Start the server
	log.Println("HTTP server started on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func addClient(hub *hub.Hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade get to a websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Create new client and run on the central relay
		client := client.CreateClient(ws)
		go client.Run(hub)
	}
}
