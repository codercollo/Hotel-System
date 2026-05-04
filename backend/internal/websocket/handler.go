package websocket

import (
	"net/http"

	pkgjwt "github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // CORS is handled by the middleware layer
	},
}

type Handler struct {
	hub        *Hub
	jwtManager *pkgjwt.Manager
}

func NewHandler(hub *Hub, jwtManager *pkgjwt.Manager) *Handler {
	return &Handler{hub: hub, jwtManager: jwtManager}
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		response.Error(w, pkgjwt.ErrMissingToken)
		return
	}
	claims, err := h.jwtManager.Parse(token)
	if err != nil {
		response.Error(w, err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// upgrader already wrote the error response
		return
	}

	NewClient(claims.UserID, h.hub, conn)
}

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Get("/ws", h.Connect)
}
