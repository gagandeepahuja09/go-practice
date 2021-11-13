package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"go-practice.com/chat-application/trace"
)

type room struct {
	// forward is a channel that will hold incoming messages that should be
	// forwarded to other clients
	forward chan []byte

	// join is a channel for clients wishing to join the room
	join chan *client

	// leave is a channel for client wishing to leave the room
	leave chan *client

	// clients holds all current clients in this room
	// if we were to access the map directly, it is possible that 2 goroutines running
	// concurrently might try to modify the map at the same time resulting in
	// corrupt memory or an unpredictable state.
	clients map[*client]bool

	// tracer will receive trace information of activity in the room.
	tracer trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			// forward message to the all clients
			for client := range r.clients {
				// it will only run one block of case code at a time. this is how
				// we ensure that r.clients map is modified by only one thing at a time.
				select {
				case client.send <- msg:
					// send the message
					// after this the write method of our client will pick it up and
					// send it down the socket to the browser.
					r.tracer.Trace("Sent to client")
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("Failed to send, clean up the client.")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// in order to use websockets, we must upgrade the HTTP connection using upgrader.upgrade
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// create a socket using upgrade
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	// create a client
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	// make the client join the room
	r.join <- client

	defer func() { r.leave <- client }()
	go client.write()

	client.read()
}
