package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket *websocket.Conn
	//send is channel on which message is sent
	send     chan *message
	room     *room
	userData map[string]interface{}
}

// read method allows client to read message from socket
func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		msg.AvatarURL, err = c.room.avatar.GetAvatarURL(c)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		output := outputMessage{
			Name:      msg.Name,
			Message:   msg.Message,
			Time:      msg.When.Format("01-02-2006 15:04:05"),
			AvatarURL: msg.AvatarURL,
		}
		err := c.socket.WriteJSON(output)
		if err != nil {
			return
		}
	}
}

// defer is a useful feature. when our code grows, there are many places to return
// but we don't need to add more calls to close connections, because "defer" will catch them all
