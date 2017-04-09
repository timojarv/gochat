package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/timojarv/gochat/client"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/user"
	"github.com/timojarv/gochat/webtoken"
)

type response struct {
	Err bool `json:"error"`
	Data string `json:"data"`
}

// Addclient
func addClient(hub *hub.Hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade get to a websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Create new client and run on the central relay
		client := client.New(ws)
		go client.Run(hub)
	}
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	// Only handle POST
	if req.Method != "POST" {
		j, _ := json.Marshal(response{true, "post only"})
		res.Write(j)
		return
	}

	// Parse the request body
	var data struct {
		Username, Password string
	}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)

	if data.Username == "" || data.Password == "" {
		j, _ := json.Marshal(response{true, "username or password not supplied"})
		res.Write(j)
		return
	}

	// Find user with specified name
	User, err := user.FindByUsername(data.Username)

	if err != nil && err.Error() == "not found" {
		User = user.New(data.Username, data.Password)
		err = User.Save()
	}

	// Catches 2 possible errors, convenient
	if err != nil {
		j, _ := json.Marshal(response{true, err.Error()})
		res.Write(j)
		return
	}

	if !User.Authenticate(data.Password) {
		j, _ := json.Marshal(response{true, "incorrect password"})
		res.Write(j)
		return
	}

	// Build and send the signed authentication token
	token, err := webtoken.New(User.Id.Hex())
	if err != nil {
		j, _ := json.Marshal(response{true, err.Error()})
		res.Write(j)
		return
	}

	j, _ := json.Marshal(response{false, token})
	res.Write(j)
}

func handleValidate(res http.ResponseWriter, req *http.Request) {
	// Only handle POST
	if req.Method != "POST" {
		j, _ := json.Marshal(response{true, "post only"})
		res.Write(j)
		return
	}

	// Parse the request body
	var data struct {
		Token string
	}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)

	token := data.Token
	if token == "" {
		j, _ := json.Marshal(response{true, "token not supplied"})
		res.Write(j)
		return
	}

	_, err = webtoken.GetSubject(token)
	if err != nil {
		j, _ := json.Marshal(response{true, err.Error()})
		res.Write(j)
		return
	}

	j, _ := json.Marshal(response{false, "ok"})
	res.Write(j)
}