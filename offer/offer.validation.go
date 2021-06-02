package offer

import (
	"unicode/utf8"
	"github.com/Jm-Zion/troc-bike-go/utils"
)


func (m Offer) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.AuthorID != nil && *m.AuthorID == 0 {
		errors[Columns.Offer.AuthorID] = utils.ErrEmptyValue
	}

	if m.Title != nil && utf8.RuneCountInString(*m.Title) > 300 {
		errors[Columns.Offer.Title] = utils.ErrMaxLength
	}

	if m.Description != nil && utf8.RuneCountInString(*m.Description) > 600 {
		errors[Columns.Offer.Description] = utils.ErrMaxLength
	}

	if m.Category != nil && *m.Category == 0 {
		errors[Columns.Offer.Category] = utils.ErrEmptyValue
	}

	if m.Design != nil && *m.Design == 0 {
		errors[Columns.Offer.Design] = utils.ErrEmptyValue
	}

	if m.LocationID != nil && *m.LocationID == 0 {
		errors[Columns.Offer.LocationID] = utils.ErrEmptyValue
	}

	if m.Size != nil && utf8.RuneCountInString(*m.Size) > 2 {
		errors[Columns.Offer.Size] = utils.ErrMaxLength
	}

	return errors, len(errors) == 0
}

func (m OfferDesign) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Theme != nil && utf8.RuneCountInString(*m.Theme) > 1000 {
		errors[Columns.OfferDesign.Theme] = utils.ErrMaxLength
	}

	if m.Color != nil && utf8.RuneCountInString(*m.Color) > 300 {
		errors[Columns.OfferDesign.Color] = utils.ErrMaxLength
	}

	return errors, len(errors) == 0
}

func (m OfferGear) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.OfferID != nil && *m.OfferID == 0 {
		errors[Columns.OfferGear.OfferID] = utils.ErrEmptyValue
	}

	if m.GearID != nil && *m.GearID == 0 {
		errors[Columns.OfferGear.GearID] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}

func (m OfferMedia) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.OfferID != nil && *m.OfferID == 0 {
		errors[Columns.OfferMedia.OfferID] = utils.ErrEmptyValue
	}

	if m.MediaID != nil && *m.MediaID == 0 {
		errors[Columns.OfferMedia.MediaID] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}

func (m BikeOffer) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}

	if m.Size != nil && utf8.RuneCountInString(*m.Size) > 2 {
		errors[Columns.BikeOffer.Size] = utils.ErrMaxLength
	}

	if m.WheelSize != nil && utf8.RuneCountInString(*m.WheelSize) > 10 {
		errors[Columns.BikeOffer.WheelSize] = utils.ErrMaxLength
	}

	if m.OfferID != nil && *m.OfferID == 0 {
		errors[Columns.BikeOffer.OfferID] = utils.ErrEmptyValue
	}

	return errors, len(errors) == 0
}
