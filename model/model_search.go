//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

const condition = "?.? = ?"

// base filters
type applier func(query *orm.Query) (*orm.Query, error)

type search struct {
	appliers []applier
}

func (s *search) apply(query *orm.Query) {
	for _, applier := range s.appliers {
		query.Apply(applier)
	}
}

func (s *search) where(query *orm.Query, table, field string, value interface{}) {

	query.Where(condition, pg.Ident(table), pg.Ident(field), value)

}

func (s *search) WithApply(a applier) {
	if s.appliers == nil {
		s.appliers = []applier{}
	}
	s.appliers = append(s.appliers, a)
}

func (s *search) With(condition string, params ...interface{}) {
	s.WithApply(func(query *orm.Query) (*orm.Query, error) {
		return query.Where(condition, params...), nil
	})
}

// Searcher is interface for every generated filter
type Searcher interface {
	Apply(query *orm.Query) *orm.Query
	Q() applier

	With(condition string, params ...interface{})
	WithApply(a applier)
}

type AccountSearch struct {
	search

	ID   *int
	Type *string
}

func (s *AccountSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Account.Alias, Columns.Account.ID, s.ID)
	}
	if s.Type != nil {
		s.where(query, Tables.Account.Alias, Columns.Account.Type, s.Type)
	}

	s.apply(query)

	return query
}

func (s *AccountSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type BikeOfferSearch struct {
	search

	ID             *int
	Size           *string
	WheelSize      *string
	OfferID        *int
	ElectricAssist *bool
}

func (s *BikeOfferSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.BikeOffer.Alias, Columns.BikeOffer.ID, s.ID)
	}
	if s.Size != nil {
		s.where(query, Tables.BikeOffer.Alias, Columns.BikeOffer.Size, s.Size)
	}
	if s.WheelSize != nil {
		s.where(query, Tables.BikeOffer.Alias, Columns.BikeOffer.WheelSize, s.WheelSize)
	}
	if s.OfferID != nil {
		s.where(query, Tables.BikeOffer.Alias, Columns.BikeOffer.OfferID, s.OfferID)
	}
	if s.ElectricAssist != nil {
		s.where(query, Tables.BikeOffer.Alias, Columns.BikeOffer.ElectricAssist, s.ElectricAssist)
	}

	s.apply(query)

	return query
}

func (s *BikeOfferSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type CategorySearch struct {
	search

	ID          *int
	Title       *string
	Description *string
	ParentID    *int
}

func (s *CategorySearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Category.Alias, Columns.Category.ID, s.ID)
	}
	if s.Title != nil {
		s.where(query, Tables.Category.Alias, Columns.Category.Title, s.Title)
	}
	if s.Description != nil {
		s.where(query, Tables.Category.Alias, Columns.Category.Description, s.Description)
	}
	if s.ParentID != nil {
		s.where(query, Tables.Category.Alias, Columns.Category.ParentID, s.ParentID)
	}

	s.apply(query)

	return query
}

func (s *CategorySearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type GearSearch struct {
	search

	ID          *int
	Title       *string
	Description *string
	Category    *int
}

func (s *GearSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Gear.Alias, Columns.Gear.ID, s.ID)
	}
	if s.Title != nil {
		s.where(query, Tables.Gear.Alias, Columns.Gear.Title, s.Title)
	}
	if s.Description != nil {
		s.where(query, Tables.Gear.Alias, Columns.Gear.Description, s.Description)
	}
	if s.Category != nil {
		s.where(query, Tables.Gear.Alias, Columns.Gear.Category, s.Category)
	}

	s.apply(query)

	return query
}

func (s *GearSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type LocationSearch struct {
	search

	ID      *int
	City    *string
	Address *string
	Zip     *string
	Country *string
}

func (s *LocationSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Location.Alias, Columns.Location.ID, s.ID)
	}
	if s.City != nil {
		s.where(query, Tables.Location.Alias, Columns.Location.City, s.City)
	}
	if s.Address != nil {
		s.where(query, Tables.Location.Alias, Columns.Location.Address, s.Address)
	}
	if s.Zip != nil {
		s.where(query, Tables.Location.Alias, Columns.Location.Zip, s.Zip)
	}
	if s.Country != nil {
		s.where(query, Tables.Location.Alias, Columns.Location.Country, s.Country)
	}

	s.apply(query)

	return query
}

func (s *LocationSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type LoginSearch struct {
	search

	ID       *int
	Login    *string
	Password *string
	UserID   *int
}

func (s *LoginSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Login.Alias, Columns.Login.ID, s.ID)
	}
	if s.Login != nil {
		s.where(query, Tables.Login.Alias, Columns.Login.Login, s.Login)
	}
	if s.Password != nil {
		s.where(query, Tables.Login.Alias, Columns.Login.Password, s.Password)
	}
	if s.UserID != nil {
		s.where(query, Tables.Login.Alias, Columns.Login.UserID, s.UserID)
	}

	s.apply(query)

	return query
}

func (s *LoginSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type MediaSearch struct {
	search

	ID        *int
	Thumbnail *string
	Raw       *string
}

func (s *MediaSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Media.Alias, Columns.Media.ID, s.ID)
	}
	if s.Thumbnail != nil {
		s.where(query, Tables.Media.Alias, Columns.Media.Thumbnail, s.Thumbnail)
	}
	if s.Raw != nil {
		s.where(query, Tables.Media.Alias, Columns.Media.Raw, s.Raw)
	}

	s.apply(query)

	return query
}

func (s *MediaSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type OfferSearch struct {
	search

	ID          *int
	Author      *int
	Title       *string
	Description *string
	Price       *int
	Negociation *bool
	Category    *int
	Design      *int
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	CreatedAt   *time.Time
	LocationID  *int
	WheelSize   *int
	Size        *string
	Condition   *int
	Enabled     *bool
	Quantity    *int
}

func (s *OfferSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.ID, s.ID)
	}
	if s.Author != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Author, s.Author)
	}
	if s.Title != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Title, s.Title)
	}
	if s.Description != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Description, s.Description)
	}
	if s.Price != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Price, s.Price)
	}
	if s.Negociation != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Negociation, s.Negociation)
	}
	if s.Category != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Category, s.Category)
	}
	if s.Design != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Design, s.Design)
	}
	if s.UpdatedAt != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.UpdatedAt, s.UpdatedAt)
	}
	if s.DeletedAt != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.DeletedAt, s.DeletedAt)
	}
	if s.CreatedAt != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.CreatedAt, s.CreatedAt)
	}
	if s.LocationID != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.LocationID, s.LocationID)
	}
	if s.WheelSize != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.WheelSize, s.WheelSize)
	}
	if s.Size != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Size, s.Size)
	}
	if s.Condition != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Condition, s.Condition)
	}
	if s.Enabled != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Enabled, s.Enabled)
	}
	if s.Quantity != nil {
		s.where(query, Tables.Offer.Alias, Columns.Offer.Quantity, s.Quantity)
	}

	s.apply(query)

	return query
}

func (s *OfferSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type OfferDesignSearch struct {
	search

	ID         *int
	Theme *string
	Color      *string
}

func (s *OfferDesignSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.OfferDesign.Alias, Columns.OfferDesign.ID, s.ID)
	}
	if s.Theme != nil {
		s.where(query, Tables.OfferDesign.Alias, Columns.OfferDesign.Theme, s.Theme)
	}
	if s.Color != nil {
		s.where(query, Tables.OfferDesign.Alias, Columns.OfferDesign.Color, s.Color)
	}

	s.apply(query)

	return query
}

func (s *OfferDesignSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type OfferGearSearch struct {
	search

	ID      *int
	OfferID *int
	GearID  *int
}

func (s *OfferGearSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.OfferGear.Alias, Columns.OfferGear.ID, s.ID)
	}
	if s.OfferID != nil {
		s.where(query, Tables.OfferGear.Alias, Columns.OfferGear.OfferID, s.OfferID)
	}
	if s.GearID != nil {
		s.where(query, Tables.OfferGear.Alias, Columns.OfferGear.GearID, s.GearID)
	}

	s.apply(query)

	return query
}

func (s *OfferGearSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type OfferMediaSearch struct {
	search

	ID      *int
	OfferID *int
	MediaID *int
}

func (s *OfferMediaSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.OfferMedia.Alias, Columns.OfferMedia.ID, s.ID)
	}
	if s.OfferID != nil {
		s.where(query, Tables.OfferMedia.Alias, Columns.OfferMedia.OfferID, s.OfferID)
	}
	if s.MediaID != nil {
		s.where(query, Tables.OfferMedia.Alias, Columns.OfferMedia.MediaID, s.MediaID)
	}

	s.apply(query)

	return query
}

func (s *OfferMediaSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type UserSearch struct {
	search

	ID          *int
	Firstname   *string
	Lastname    *string
	MobilePhone *string
	Email       *string
	AccountID   *int
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	CreatedAt   *time.Time
}

func (s *UserSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.User.Alias, Columns.User.ID, s.ID)
	}
	if s.Firstname != nil {
		s.where(query, Tables.User.Alias, Columns.User.Firstname, s.Firstname)
	}
	if s.Lastname != nil {
		s.where(query, Tables.User.Alias, Columns.User.Lastname, s.Lastname)
	}
	if s.MobilePhone != nil {
		s.where(query, Tables.User.Alias, Columns.User.MobilePhone, s.MobilePhone)
	}
	if s.Email != nil {
		s.where(query, Tables.User.Alias, Columns.User.Email, s.Email)
	}
	if s.AccountID != nil {
		s.where(query, Tables.User.Alias, Columns.User.AccountID, s.AccountID)
	}
	if s.UpdatedAt != nil {
		s.where(query, Tables.User.Alias, Columns.User.UpdatedAt, s.UpdatedAt)
	}
	if s.DeletedAt != nil {
		s.where(query, Tables.User.Alias, Columns.User.DeletedAt, s.DeletedAt)
	}
	if s.CreatedAt != nil {
		s.where(query, Tables.User.Alias, Columns.User.CreatedAt, s.CreatedAt)
	}

	s.apply(query)

	return query
}

func (s *UserSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}
