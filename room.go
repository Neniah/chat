package main

type room struct {
	// forward is a channel that holds incoming messages
	// that should be fowwarded to the other clients.
	forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool
}


