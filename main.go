package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/timojarv/gochat/gopher"
	"github.com/timojarv/gochat/hub"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/validate", handleValidate)

	// Init a new message relay
	hub := hub.New()

	// Websocket route adds new client to the relay
	http.HandleFunc("/ws", addClient(hub))

	// Start the message relay
	go hub.Run()

	// Initialize the @gopher bot
	gopher.Register(hub)

	// Start the server
	log.Println("HTTP server started on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}