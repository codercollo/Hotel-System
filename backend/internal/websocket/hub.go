package websocket

import (
	"encoding/json"
	"sync"

	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// Hub maintains the set of active clients and broadcasts messages.
type Hub struct {
	mu      sync.RWMutex
	clients map[string]map[*Client]struct{} // userID → set of clients
}

// NewHub creates a ready-to-use Hub.
func NewHub() *Hub {
	return &Hub{clients: make(map[string]map[*Client]struct{})}
}

// Register adds a client to the hub.
func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[c.UserID]; !ok {
		h.clients[c.UserID] = make(map[*Client]struct{})
	}
	h.clients[c.UserID][c] = struct{}{}
}

// Unregister removes a client from the hub.
func (h *Hub) Unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if clients, ok := h.clients[c.UserID]; ok {
		delete(clients, c)
		if len(clients) == 0 {
			delete(h.clients, c.UserID)
		}
	}
}

// SendToUser delivers a message to all connections belonging to userID.
func (h *Hub) SendToUser(userID string, msg *Message) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	clients, ok := h.clients[userID]
	if !ok {
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		logger.Error(err, "ws hub: marshal message")
		return
	}

	for c := range clients {
		select {
		case c.send <- b:
		default:
			logger.Warn("ws hub: client send buffer full, dropping message")
		}
	}
}

// Broadcast sends a message to every connected client.
func (h *Hub) Broadcast(msg *Message) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	for _, clients := range h.clients {
		for c := range clients {
			select {
			case c.send <- b:
			default:
			}
		}
	}
}
