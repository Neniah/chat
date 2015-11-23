package main

type room struct {
	// forward is a channel that holds incoming messages
	// that should be fowwarded to the other clients.
	forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool
}

func (r *room) run(){
	for{
		select {
		case client := <-r.join:
			// Joining
			r.clients[client] = true
		case client := <-r.leave:
			// Leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <--r.forward:
			// forward message to call client
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send the message
				default:
					//failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

