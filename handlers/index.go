package handlers

import (
	"articles-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := storage.GetAllArticles()

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, articles)
	case "application/xml":
		c.XML(http.StatusOK, articles)
	default:
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"articles": articles,
			},
		)
	}
}
