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
	// CreateSchema(db.GetDB())

	var n New

	n.Name = "Ratinho espanca goku"
	n.CategoryName = "Notícia"
	n.Description = "Ratinho espanca goku ao vivo"
	n.ImageUrl = "https://imagem.com/imagem.ong"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: 1, Content: "OLA", ImageUrl: "https://imagem.com/png"}

	// fmt.Println(n.Rows[0])
	// fmt.Println("aqui")

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

func TestSearchNewWithRows(t *testing.T) {
	config.Init()
	db.Init()

	new, err2 := SearchNewWithRows(6)

	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(new)
}
