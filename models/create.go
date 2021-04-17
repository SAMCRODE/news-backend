package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

func CreateSchema(pg *pg.DB) {
	qs := []string{
		"CREATE TABLE news (id int, name text, description text, category_name text, image_url text)",
		"CREATE TABLE rows (id int, type int, content text, imageUrl text, new_id int)",
		"INSERT into news VALUES (1, 'macaco bate em godzila', 'noticia top', 'ultimas', 'macaco.com.br/png')",
		"INSERT into rows VALUES (1, 1, '<p>noticia</p>', 'foto.png', 1)",
	}
	for _, q := range qs {
		fmt.Println(q)
		_, err := pg.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}
