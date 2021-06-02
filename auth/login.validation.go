package auth

import (
	"unicode/utf8"
	model "github.com/Jm-Zion/troc-bike-go/model"
	"github.com/Jm-Zion/troc-bike-go/utils"
)

func (m Login) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if utf8.RuneCountInString(m.Login) > 300 {
		errors[model.Columns.Login.Login] = utils.ErrMaxLength
	}

	if utf8.RuneCountInString(m.Password) > 300 {
		errors[model.Columns.Login.Password] = utils.ErrMaxLength
	}

	if m.UserID != nil && *m.UserID == 0 {
		errors[model.Columns.Login.UserID] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}