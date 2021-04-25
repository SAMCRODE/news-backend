package models

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/news-backend/db"
)

type RowType string

const (
	ROWLAYOUT1TYPE = "ROWLAYOUT1TYPE"
	ROWLAYOUT2TYPE = "ROWLAYOUT2TYPE"
	ROWPUREHTML    = "ROWPUREHTML"
)

type Row struct {
	tableName    struct{} `pg:"rows"`
	Id           int      `pg:",pk"`
	Type         RowType
	Content      string
	ImageUrl     string
	NewId        int
	FirstLetter  string
	CaptionImage string
}

type New struct {
	tableName       struct{} `pg:"news"`
	Id              int      `pg:",pk"`
	Name            string
	Description     string
	CategoryName    string
	ImageUrl        string
	BackgroundColor string
	ShowNumber      int
	CreateDate      time.Time
	Rows            []*Row `pg:"rel:has-many"`
}

func (n New) Save() error {
	pg := db.GetDB()
	_, err := pg.Model(&n).Returning("Id").Insert()

	for i := 0; i < len(n.Rows); i++ {
		n.Rows[i].NewId = n.Id
		_, err := pg.Model(n.Rows[i]).Insert()
		fmt.Println(err)
	}

	return err
}

func SearchNews() ([]New, error) {
	pg := db.GetDB()
	var news []New

	err := pg.Model(&news).Select()

	return news, err
}

func SearchNewsPaginated(page int) ([]New, error) {
	pg := db.GetDB()
	var news []New

	err := pg.Model(&news).Offset(10 * (page - 1)).Limit(10).Select()

	return news, err
}

func SearchLastestNews(cnt int) ([]New, error) {
	pg := db.GetDB()
	var news []New

	err := pg.Model(&news).OrderExpr("id DESC").Where("show_number = 0").Limit(cnt).Select()

	return news, err
}

func SearchHotNews() ([]New, error) {
	pg := db.GetDB()
	var news []New

	err := pg.Model(&news).Where("show_number > 0").Select()

	return news, err
}

func SearchNewWithRows(id int) (New, error) {
	db := db.GetDB()
	var new New
	err := db.Model(&new).Column("new.*").Where("? = ?", pg.Ident("id"), id).Relation("Rows").First()

	return new, err
}
