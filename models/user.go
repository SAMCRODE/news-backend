package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/news-backend/db"
)

type User struct {
	tableName   struct{} `pg:"users"`
	Id          int      `pg:",pk"`
	Name        string
	Email       string
	Password    string
	ImageUrl    string
	Permissions int
}

func (n User) Save() error {
	pg := db.GetDB()
	hash := md5.Sum([]byte(n.Password))
	n.Password = hex.EncodeToString(hash[:])

	_, err := pg.Model(&n).Returning("Id").Insert()

	return err
}

func SearchUserByEmail(email string) (User, error) {
	pg := db.GetDB()
	var user User
	_, err := pg.Query(&user, `SELECT * FROM users WHERE email = ?`, email)

	return user, err
}
