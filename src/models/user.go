package models

import (
	"errors"

	"github.com/tejasvi541/Go-Server/src/db"
	"github.com/tejasvi541/Go-Server/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {

	query :=`
		SELECT id FROM users WHERE email = $1
	`
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID)
	if err == nil {
		return errors.New("user already exists")
	}

	query = `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	hashedPw, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.Username, u.Email, hashedPw).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Authenticate() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	var hashedPw string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &hashedPw)
	if err != nil {
		return errors.New("invalid email or password")
	}

	if err := utils.CheckPassword(u.Password, hashedPw); err != nil {
		return errors.New("invalid email or password")
	}

	return nil
}



