package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// Message represents the format of incoming messages from the client.
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Client represents a connected WebSocket user.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	room string
}

// readPump listens for incoming WebSocket messages and routes them to the hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, rawMsg, err := c.conn.ReadMessage()
		if err != nil {
			break // client disconnected
		}

		var msg Message
		if err := json.Unmarshal(rawMsg, &msg); err != nil {
			continue // ignore malformed input
		}

		formatted := []byte("[" + msg.Username + "]: " + msg.Message)

		c.hub.broadcast <- RoomMessage{
			room:    c.room,
			message: formatted,
		}
	}
}

// writePump sends messages from the hub to the client's WebSocket.
func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// Channel closed, close socket
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}
