package gear

import (
	cat "github.com/Jm-Zion/troc-bike-go/category"
)

type Gear struct {
	tableName struct{} `sql:"gear,alias:t" pg:",discard_unknown_columns"`

	ID          int     `sql:"id,pk"`
	Title       *string `sql:"title"`
	Description *string `sql:"description"`
	Category    *int    `sql:"category"`

	CategoryRel  *cat.Category `pg:"fk:category"`
	CategoryRel1 *cat.Category `pg:"fk:category"`
}

var Columns = struct {
	Gear struct {
		ID, Title, Description, Category string

		CategoryRel, CategoryRel1 string
	}
}{
	Gear: struct {
		ID, Title, Description, Category string

		CategoryRel, CategoryRel1 string
	}{
		ID:          "id",
		Title:       "title",
		Description: "description",
		Category:    "category",

		CategoryRel:  "CategoryRel",
		CategoryRel1: "CategoryRel1",
	},
}

var Tables = struct {
	Gear struct {
		Name, Alias string
	}
}{
	Gear: struct {
		Name, Alias string
	}{
		Name:  "gear",
		Alias: "t",
	},
}