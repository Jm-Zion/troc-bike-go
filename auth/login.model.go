package auth

import (
	 "github.com/Jm-Zion/troc-bike-go/user"
)

type Login struct {
	tableName struct{} `sql:"login" pg:"login"`

	ID       int    `sql:"id,pk"`
	Login    string `sql:"login,notnull" json:"identifier" validate:"required"`
	Password string `sql:"password,notnull" json:"password" validate:"required"`
	UserID   *int   `sql:"user_id"`

	User *user.User `pg:"fk:user_id"`
}
