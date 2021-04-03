package models

import "github.com/news-backend/db"

type New struct {
	ID           int
	Name         string
	Description  string
	CategoryName string
	ImageUrl     string
}

func (n New) Save() error {
	pg := db.GetDB()
	_, err := pg.Model(&n).Returning("ID").Insert()

	return err
}

func SearchNews() ([]New, error) {
	pg := db.GetDB()
	var news []New

	err := pg.Model(&news).Select()

	return news, err
}
