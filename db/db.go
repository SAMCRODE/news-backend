package db

import (
	"github.com/news-backend/config"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func Init() {
	conf := config.GetConfig()
	opt, err := pg.ParseURL(conf.DATABASEURI)

	if err == nil {
		db = pg.Connect(opt)
		err = db.Ping(db.Context())
	}

	if err != nil {
		panic(err)
	}
}

func GetDB() *pg.DB {
	return db
}
