package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/news-backend/config"
)

func CreateSchema(pg *pg.DB) {
	temporary := config.GetConfig().MODE != "production"

	for _, model := range []interface{}{&User{}, &New{}, &Row{}} {
		err := pg.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:          temporary, // create temp table if not on production
			FKConstraints: true,
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}
