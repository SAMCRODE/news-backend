package models

import (
	"fmt"
	"os"
	"testing"

	"github.com/news-backend/config"
	"github.com/news-backend/db"
)

func TestMain(m *testing.M) {
	config.Init()
	db.Init()
	os.Exit(m.Run())
}

func TestInsert(t *testing.T) {
	// CreateSchema(db.GetDB())
	fmt.Println("Aqui")
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
	news, err2 := SearchNews()

	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(news)
}

func TestSearchNewWithRows(t *testing.T) {
	new, err2 := SearchNewWithRows(1)

	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(new)
}

func TestSearchNewsPaginated(t *testing.T) {
	new, err := SearchNewsPaginated(1)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(new)
}
