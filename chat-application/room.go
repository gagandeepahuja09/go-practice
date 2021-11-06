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
