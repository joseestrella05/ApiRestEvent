package models

import (
	"example.com/rest-api/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.password)

	if err != nil {
		return err
	}

	UserId, err := result.LastInsertId()

	u.ID = UserId

	return err
}
