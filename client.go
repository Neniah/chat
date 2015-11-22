package main

import (
	"github.com/gorilla/websocket"
)

//client representes a single chatting user
type client struct{
	//socket is the web socket for this client.
	socket *websocket.Conn
	// send in a channel on which messages are sent.
	send chan []byte
	// room is the room this client is chatting in.
	room *room
}

