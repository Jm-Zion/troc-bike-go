package user

import (
	"unicode/utf8"
	"regexp"
	"github.com/Jm-Zion/troc-bike-go/utils"
)

var Columns = struct {
	Account struct {
		ID, Type string
	}
	User struct {
		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

		Account string
	}
}{
	Account: struct {
		ID, Type string
	}{
		ID:   "id",
		Type: "type",
	},
	User: struct {
		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

		Account string
	}{
		ID:          "id",
		Firstname:   "firstname",
		Lastname:    "lastname",
		MobilePhone: "mobile_phone",
		Email:       "email",
		AccountID:   "account_id",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		CreatedAt:   "created_at",

		Account: "Account",
	},
}

func (m Account) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Type != nil && utf8.RuneCountInString(*m.Type) > 300 {
		errors[Columns.Account.Type] = utils.ErrMaxLength
	}

	return errors, len(errors) == 0
}


func (m User) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Firstname != nil && utf8.RuneCountInString(*m.Firstname) > 300 {
		errors[Columns.User.Firstname] = utils.ErrMaxLength
	}

	if m.Lastname != nil && utf8.RuneCountInString(*m.Lastname) > 300 {
		errors[Columns.User.Lastname] = utils.ErrMaxLength
	}

	re := regexp.MustCompile(`^(?:(?:\+|00)33[\s.-]{0,3}(?:\(0\)[\s.-]{0,3})?|0)[1-9](?:(?:[\s.-]?\d{2}){4}|\d{2}(?:[\s.-]?\d{3}){2})$`)

	if (m.MobilePhone != nil && utf8.RuneCountInString(*m.MobilePhone) > 300) || (re.MatchString(*m.MobilePhone) != true) {
		errors[Columns.User.MobilePhone] = utils.ErrMaxLength
	}

	if m.Email != nil && utf8.RuneCountInString(*m.Email) > 300 {
		errors[Columns.User.Email] = utils.ErrMaxLength
	}

	if m.AccountID != nil && *m.AccountID == 0 {
		errors[Columns.User.AccountID] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}
