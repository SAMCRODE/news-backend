package main

import (
	"github.com/news-backend/config"
	"github.com/news-backend/db"
	"github.com/news-backend/models"
)

func main() {
	config.Init()
	db.Init()

	models.CreateSchema(db.GetDB())
}
