package models

import (
	"encoding/json"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func (m *ClientManager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.Clients[client] = true
			jsonMessage, err := json.Marshal(&Message{Content: "/A new socket has connected."})
			if err != nil {
				m.send(jsonMessage, client)
			}
		case client := <-m.Unregister:
			if _, ok := m.Clients[client]; ok {
				close(client.Send)
				delete(m.Clients, client)
				jsonMessage, err := json.Marshal(&Message{Content: "/A socket has disconnected."})
				if err != nil {
					m.send(jsonMessage, client)
				}
			}
		case message := <-m.Broadcast:
			for client := range m.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(m.Clients, client)
				}
			}
		}
	}
}

func (m *ClientManager) send(jsonMessage []byte, client *Client) {
	for currClient := range m.Clients {
		if currClient != client {
			currClient.Send <- jsonMessage
		}
	}
}
