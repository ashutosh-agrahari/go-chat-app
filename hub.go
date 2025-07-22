package main

// RoomMessage contains the room name and the message to be broadcast.
type RoomMessage struct {
	room    string
	message []byte
}

// Hub manages all clients and routes messages to the appropriate room.
type Hub struct {
	clients    map[string]map[*Client]bool // room → set of clients
	broadcast  chan RoomMessage
	register   chan *Client
	unregister chan *Client
}

// newHub initializes and returns a new Hub instance.
func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		broadcast:  make(chan RoomMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// run handles registration, unregistration, and message broadcasting for each room.
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if h.clients[client.room] == nil {
				h.clients[client.room] = make(map[*Client]bool)
			}
			h.clients[client.room][client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client.room][client]; ok {
				delete(h.clients[client.room], client)
				close(client.send)
				if len(h.clients[client.room]) == 0 {
					delete(h.clients, client.room)
				}
			}

		case rm := <-h.broadcast:
			if clientsInRoom, ok := h.clients[rm.room]; ok {
				for client := range clientsInRoom {
					select {
					case client.send <- rm.message:
					default:
						close(client.send)
						delete(clientsInRoom, client)
					}
				}
			}
		}
	}
}
