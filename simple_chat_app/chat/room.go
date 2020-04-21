package main

import (
	"github.com/stretchr/objx"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan *message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	avatar  Avatar
}

func NewRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		avatar:  avatar,
	}
}

// for keyword: loop forever
// select...case: it will only run one block of case code at a time
// this is how we are able to synchronize to ensure that our r.clients map
// is only ever modified by one thing at a time
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP", err)
		return
	}
	UserData := map[string]interface{}{}

	if authCookie, err := req.Cookie("auth"); err == nil {
		UserData = objx.MustFromBase64(authCookie.Value)
	}

	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: UserData,
	}
	r.join <- client
	defer func() { r.leave <- client }()

	//inside read() we have "for" loop forever => it keeps connection alive
	// we can also switch like this
	// go client.read()
	// client.write()
	go client.write()
	client.read()

}
