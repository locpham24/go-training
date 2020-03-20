package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	//send is channel on which message is sent
	send chan []byte
	room *room
}

// read method allows client to read message from socket
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

// defer is a useful feature. when our code grows, there are many places to return
// but we don't need to add more calls to close connections, because "defer" will catch them all
