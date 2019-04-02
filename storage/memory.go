package storage

import (
	"articles-service/models"
	"errors"
)

var articleList = []models.Article{
	models.Article{ID: "1", Title: "Title 1", Content: "Content 1"},
	models.Article{ID: "2", Title: "Title 2", Content: "Content 2"},
}

func GetAllArticles() []models.Article {
	return articleList
}

func GetArticleByID(id string) (*models.Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
