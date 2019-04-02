package handlers

import (
	"articles-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ShowWebSocketPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"websocket.html",
		gin.H{},
	)
}

func WS(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	manager := c.MustGet("manager").(*models.ClientManager)

	client := models.NewClient(id.String(), conn, manager)

	manager.Register <- client

	go client.Read()
	go client.Write()
}
