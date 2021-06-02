package location

import (
	"unicode/utf8"
	"github.com/Jm-Zion/troc-bike-go/utils"
)

func (m Location) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.City != nil && utf8.RuneCountInString(*m.City) > 300 {
		errors[Columns.Location.City] = utils.ErrMaxLength
	}

	if m.Address != nil && utf8.RuneCountInString(*m.Address) > 300 {
		errors[Columns.Location.Address] = utils.ErrMaxLength
	}

	if m.Zip != nil && utf8.RuneCountInString(*m.Zip) > 300 {
		errors[Columns.Location.Zip] = utils.ErrMaxLength
	}

	if m.Country != nil && utf8.RuneCountInString(*m.Country) > 300 {
		errors[Columns.Location.Country] = utils.ErrMaxLength
	}

	return errors, len(errors) == 0
}
