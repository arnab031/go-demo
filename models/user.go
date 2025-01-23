package models

import (
	"errors"
	"fmt"

	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err

}

func (u *User) Authenticate() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var id int64
	var hashedPassword string
	err := row.Scan(&id, &hashedPassword)
	fmt.Println(u.Email, u.Password, hashedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, hashedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

    u.ID = id
	return nil
}
