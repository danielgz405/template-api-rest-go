package websocket

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Config struct {
	Port      string
	JWTSecret string
	DbURI     string
}

type Server interface {
	Config() *Config
	Hub() *Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *Hub
}

type Client struct {
	hub      *Hub
	id       string
	roles    []string
	module   string
	socket   *websocket.Conn
	outbound chan []byte
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (c *Client) Write() {
	for {
		select {
		case message, ok := <-c.outbound:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
