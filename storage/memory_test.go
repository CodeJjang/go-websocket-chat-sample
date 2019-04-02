package storage

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	articles := GetAllArticles()

	if len(articles) != len(articleList) {
		t.Fail()
	}

	for idx, article := range articles {
		if article != articleList[idx] {
			t.Fail()
		}
	}
}