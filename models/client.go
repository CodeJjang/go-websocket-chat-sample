package models

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
	manager *ClientManager
}

func NewClient(id string, socket *websocket.Conn, manager *ClientManager) *Client {
	return &Client{
		ID:     id,
		Socket: socket,
		Send:   make(chan []byte),
		manager: manager,
	}
}

func (c *Client) Read() {
	defer func() {
		// c.manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			c.manager.Unregister <- c
			c.Socket.Close()
			return
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		c.manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		// c.manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <- c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
                return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
