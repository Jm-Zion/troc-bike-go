package category

import (
	"unicode/utf8"
	"github.com/Jm-Zion/troc-bike-go/utils"
)

func (m Category) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Title != nil && utf8.RuneCountInString(*m.Title) > 300 {
		errors[Columns.Category.Title] = utils.ErrMaxLength
	}

	if m.Description != nil && utf8.RuneCountInString(*m.Description) > 1000 {
		errors[Columns.Category.Description] = utils.ErrMaxLength
	}

	return errors, len(errors) == 0
}
