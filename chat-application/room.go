package main

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
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)

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
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
