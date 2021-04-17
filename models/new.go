package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/news-backend/db"
)

type Row struct {
	tableName struct{} `pg:"rows"`
	Id        int      `pg:",pk"`
	Type      int
	Content   string
	ImageUrl  string
	NewId     int
}

type New struct {
	tableName    struct{} `pg:"news"`
	Id           int      `pg:",pk"`
	Name         string
	Description  string
	CategoryName string
	ImageUrl     string
	Rows         []*Row `pg:"rel:has-many"`
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

func SearchNewWithRows(id int) (New, error) {
	db := db.GetDB()
	var new New
	err := db.Model(&new).Column("new.*").Where("? = ?", pg.Ident("id"), id).Relation("Rows").First()

	return new, err
}
