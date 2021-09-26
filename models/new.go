package models

import (
	"time"

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
	AuthorId        int
	Author          User `pg:"rel:has-one"`
}

func (n *New) AdjustRows() error {
	pg := db.GetDB()
	mnew, err := SearchNewWithRows(n.Id)

	if err != nil {
		return err
	}

	for i := 0; i < len(mnew.Rows); i++ {
		found := false

		for j := 0; j < len(n.Rows); j++ {
			if n.Rows[j].Id == mnew.Rows[i].Id {
				found = true
				break
			}
		}

		if !found {
			_, err := pg.Model(mnew.Rows[i]).WherePK().ForceDelete()

			if err != nil {
				return err
			}
		}
	}

	return err
}

func (n *New) Save() error {
	pg := db.GetDB()
	var err error

	if n.Id == 0 {
		_, err = pg.Model(n).Returning("Id").Insert()
	} else {
		_, err = pg.Model(n).Returning("Id").Where("new.id = ?", n.Id).Update()
	}

	for i := 0; i < len(n.Rows); i++ {
		n.Rows[i].NewId = n.Id

		if n.Rows[i].Id == 0 {
			pg.Model(n.Rows[i]).Insert()
		} else {
			_, err = pg.Model(n.Rows[i]).Returning("Id").Where("row.id = ?", n.Rows[i].Id).Update()
		}
	}

	if err != nil {
		return err
	}

	return n.AdjustRows()
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

func SearchNewWithRows(id int) (*New, error) {
	db := db.GetDB()
	var new New
	err := db.Model(&new).
		Relation("Author").
		Where("new.id = ?", id).
		Relation("Rows").Select()

	new.Author.Password = ""
	new.Author.Permissions = 0
	new.Author.Id = 0

	return &new, err
}
