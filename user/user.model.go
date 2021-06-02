package user

import (
	"time"
)


// var Columns = struct {
// 	Account struct {
// 		ID, Type string
// 	}
// 	User struct {
// 		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

// 		Account string
// 	}
// }{
// 	Account: struct {
// 		ID, Type string
// 	}{
// 		ID:   "id",
// 		Type: "type",
// 	},
// 	User: struct {
// 		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

// 		Account string
// 	}{
// 		ID:          "id",
// 		Firstname:   "firstname",
// 		Lastname:    "lastname",
// 		MobilePhone: "mobile_phone",
// 		Email:       "email",
// 		AccountID:   "account_id",
// 		UpdatedAt:   "updated_at",
// 		DeletedAt:   "deleted_at",
// 		CreatedAt:   "created_at",

// 		Account: "Account",
// 	},
// }

// var Tables = struct {
// 	Account struct {
// 		Name, Alias string
// 	}
// 	User struct {
// 		Name, Alias string
// 	}
// }{
// 	Account: struct {
// 		Name, Alias string
// 	}{},
// 	User: struct {
// 		Name, Alias string
// 	}{
// 		Name:  "user",
// 		Alias: "t",
// 	},
// }

type User struct {
	tableName struct{} `sql:"user" pg:"user"`

	ID          int        `sql:"id,pk" json:"id"`
	Firstname   *string    `sql:"firstname" json:"firstname"`
	Lastname    *string    `sql:"lastname" json:"lastname"`
	Username    *string    `sql:"username" json:"username"`
	MobilePhone *string    `sql:"mobile_phone" json:"mobilePhone"`
	Email       *string    `sql:"email" json:"email"`
	AccountID   *int       `sql:"account_id" json:"accountID"`
	UpdatedAt   *time.Time `sql:"updated_at" json:"updatedAt"`
	DeletedAt   *time.Time `sql:"deleted_at" json:"deletedAt"`
	CreatedAt   *time.Time `sql:"created_at" json:"createdAt"`

	Account *Account `pg:"fk:account_id"`
}

type Account struct {
	tableName struct{} `sql:"account" pg:"account"`

	ID   int     `sql:"id,pk"`
	Type *string `sql:"type"`
}