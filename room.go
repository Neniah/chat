package main

import (
	"log"
	"net/http"
)

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
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			// Leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <--r.forward:
			// forward message to call client
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send the message
					r.tracer.Trace(" -- sent to client")
				default:
					//failed to send
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" --  failed to send, cleaned up client")
				}
			}
		}
	}
}

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}
	client := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}
	r.join <- client
	defer func() { r.leave <- client } ()
	go client.write()
	client.read()
}


// newRoom makes a new room that is ready to go.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join: 	 make(chan *client),
		leave: 	 make(chan *client),
		clients: make(map[*client]book),
	}
}
