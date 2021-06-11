package models

import "testing"

func TestInsertUser(t *testing.T) {
	var u User

	u.Name = "Milionario"
	u.Email = "myemailemail@email.com"
	u.ImageUrl = "https://www.daynews.com.br/wp-content/uploads/2017/07/Jos%C3%A9-Rico-570x380.png"
	u.Password = "thatisyourlastword?"

	err := u.Save()

	if err != nil {
		t.Error(err)
	}
}
