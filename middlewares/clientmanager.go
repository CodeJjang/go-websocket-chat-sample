package middlewares

import (
	"articles-service/models"

	"github.com/gin-gonic/gin"
)

func ClientManager() gin.HandlerFunc {
	manager := &models.ClientManager{
		Clients:    make(map[*models.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
	}

	go manager.Start()
	return func(c *gin.Context) {
		c.Set("manager", manager)
		c.Next()
	}
}
