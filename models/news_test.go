package models

import (
	"fmt"
	"os"
	"testing"
	"time"

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
	n.Description = "Preço da carne ultrapassa a barra de ouro"
	n.ImageUrl = "https://imagem.com/imagem.ong"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertRowType1(t *testing.T) {
	// CreateSchema(db.GetDB())
	fmt.Println("Aqui")
	var n New

	n.Name = "Ratinho espanca goku"
	n.CategoryName = "Notícia"
	n.Description = "Imagens se espalha e choca otakus"
	n.ImageUrl = "https://imagem.com/imagem.ong"
	n.AuthorId = 1
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWLAYOUT1TYPE, Content: " melhor amigo de goku, Verdita, disse com lágrimas de ódio",
		FirstLetter:  "O",
		CaptionImage: "Melhor amigo de Goku chora",
		ImageUrl:     "https://maisdeoitomil.files.wordpress.com/2015/04/vegeta-chorando-lagrimas-negras.png?w=400"}
	n.CreateDate = time.Now()

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber1New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "Afinal, Santana tem carburador ou não tem carburador ?"
	n.CategoryName = "Automotivos"
	n.Description = "Preço da carne ultrapassa a barra de ouro"
	n.ImageUrl = "https://images.uncyc.org/pt/thumb/d/d5/Ti%C3%A3o_Carburador.jpg/1200px-Ti%C3%A3o_Carburador.jpg"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWLAYOUT1TYPE, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 1
	n.BackgroundColor = "#620607"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber2New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "Confirmada nova capa do GTA VI"
	n.CategoryName = "Games"
	n.Description = "..."
	n.ImageUrl = "https://theshoppers.com/pt-br/wp-content/uploads/sites/10/agostinho-carrara-no-gta-6-arte-de-dann1230-twitter-the-shoppers.jpg"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 2
	n.BackgroundColor = "#1A1A10"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber3New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "Ratinho espanca Goku"
	n.CategoryName = "Televisão"
	n.Description = "..."
	n.ImageUrl = "https://pbs.twimg.com/media/D9EaDlrUIAEuiVW.png:large"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 3
	n.BackgroundColor = "#CA2E30"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber4New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "Mamaco vs Godzila em breve nos cinemas"
	n.CategoryName = "Cinema"
	n.Description = "..."
	n.ImageUrl = "https://64.media.tumblr.com/f2dc10439bc1ee81fcccdb388ddabee8/fe135097f81936bb-e3/s500x750/645ad101f8753b0d490b040d017d13a100e850b4.gifv"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 4
	n.BackgroundColor = "#620607"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber5New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "O Jovem que ficou milionário com jogo do bixo"
	n.CategoryName = "Economia"
	n.Description = "..."
	n.ImageUrl = "https://blog.rico.com.vc/hs-fs/hubfs/o-que-e-ser-trader.jpg?width=607&name=o-que-e-ser-trader.jpg"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 5
	n.BackgroundColor = "#16224C"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestInsertShowNumber6New(t *testing.T) {
	// CreateSchema(db.GetDB())
	var n New

	n.Name = "Aprenda a capotar seu UNO"
	n.CategoryName = "Automotivos"
	n.Description = "..."
	n.ImageUrl = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTxnx--cqeiQi3FkjBoZduHI0P2D83xkPx79g&usqp=CAU"
	n.Rows = make([]*Row, 1)
	n.Rows[0] = &Row{Type: ROWPUREHTML, Content: "OLA", ImageUrl: "https://imagem.com/png"}
	n.CreateDate = time.Now()
	n.ShowNumber = 6
	n.BackgroundColor = "#7DC0D1"

	err2 := n.Save()

	if err2 != nil {
		panic(err2.Error())
	}
}

func TestSearch(t *testing.T) {
	_, err2 := SearchNews()

	if err2 != nil {
		panic(err2.Error())
	}

	// fmt.Println(news)
}

func TestSearchNewWithRows(t *testing.T) {
	new, err2 := SearchNewWithRows(1)

	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(new)
}

func TestSearchNewsPaginated(t *testing.T) {
	_, err := SearchNewsPaginated(1)

	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(new)
}

func TestSearchLastestNews(t *testing.T) {
	_, err := SearchLastestNews(3)

	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(news)
}

func TestSearchHotNews(t *testing.T) {
	_, err := SearchHotNews()

	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(news)
}
