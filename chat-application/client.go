package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	// socket is the websocket used for this client.
	socket *websocket.Conn
	// send is a buffered which is used to queue the received messages
	// ready to be forwarded to the client's browser
	send chan []byte
	// room is the room the client is chatting in
	room *room
}

// read method is read for the websocket, which means it will read the messages
// written by the client sent from the frontend
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			// store the message in forward channel
			c.room.forward <- msg
			// the http handler for message will read from forward
			// and send to the send channel of all clients.
		} else {
			break
		}
	}
	c.socket.Close()
}

// this will write all the messages to websocket which are ready to be
// forwarded to the client's browser. These are present in the send channel
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}