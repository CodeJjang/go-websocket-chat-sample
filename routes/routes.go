package routes

import (
	"articles-service/handlers"
	"articles-service/middlewares"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	router.Use(middlewares.ClientManager())
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	router.GET("/websocket", handlers.ShowWebSocketPage)
	router.GET("/ws", handlers.WS)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	templatesPath, err := filepath.Abs(filepath.Dir("templates/"))
	if err != nil {
		panic(err)
	}

	router.LoadHTMLGlob(path.Join(templatesPath, "*"))

	initializeRoutes(router)

	return router
}
