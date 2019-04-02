package handlers

import (
	"articles-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	articleID := c.Param("article_id")
	// Check if the article exists
	if article, err := storage.GetArticleByID(articleID); err == nil {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"article.html",
			// Pass the data that the page uses
			gin.H{
				"title":   article.Title,
				"article": article,
			},
		)

	} else {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}
}
