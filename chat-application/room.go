package main

type room struct {
	// forward is a channel that will hold incoming messages that should be
	// forwarded to other clients
	forward chan []byte
}
