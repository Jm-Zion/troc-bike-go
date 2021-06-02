package category

type Category struct {
	tableName struct{} `sql:"category" pg:"category"`

	ID          int     `sql:"id,pk" json:"id"`
	Title       *string `sql:"title" json:"title"`
	Description *string `sql:"description" json:"description"`
	ParentID    *int    `sql:"parent_id" json:"parent_id"`

	Parent *Category `pg:"fk:parent_id" json:"parent"`
}

var Columns = struct {

	Category struct {
		ID, Title, Description, ParentID string

		Parent string
	}
}{
	Category: struct {
		ID, Title, Description, ParentID string

		Parent string
	}{
		ID:          "id",
		Title:       "title",
		Description: "description",
		ParentID:    "parent_id",

		Parent: "Parent",
	},
}

var Tables = struct {
	Category struct {
		Name, Alias string
	}
}{
	Category: struct {
		Name, Alias string
	}{
		Name:  "category",
		Alias: "t",
	},
}

