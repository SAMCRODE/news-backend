package models

import (
	"fmt"
	"testing"

	"github.com/news-backend/config"
	"github.com/news-backend/db"
)

func TestInsert(t *testing.T) {
	config.Init()
	db.Init()

	var n New

	n.Name = "Ratinho espanca goku"
	n.CategoryName = "Notícia"
	n.Description = "Ratinho espanca goku ao vivo"
	n.ImageUrl = "https://imagem.com/imagem.ong"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestSearch(t *testing.T) {
	config.Init()
	db.Init()

	news, err2 := SearchNews()

	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(news)
}
