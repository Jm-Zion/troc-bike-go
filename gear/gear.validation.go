package gear

import (
	"unicode/utf8"
	"github.com/Jm-Zion/troc-bike-go/utils"
)

func (m Gear) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Title != nil && utf8.RuneCountInString(*m.Title) > 300 {
		errors[Columns.Gear.Title] = utils.ErrMaxLength
	}

	if m.Description != nil && utf8.RuneCountInString(*m.Description) > 500 {
		errors[Columns.Gear.Description] = utils.ErrMaxLength
	}

	if m.Category != nil && *m.Category == 0 {
		errors[Columns.Gear.Category] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}