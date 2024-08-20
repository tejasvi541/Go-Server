package models

import (
	"github.com/tejasvi541/Go-Server/src/db"
	"github.com/tejasvi541/Go-Server/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
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


