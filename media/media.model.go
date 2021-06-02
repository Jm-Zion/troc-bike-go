package media


type Media struct {
	tableName struct{} `sql:"media,alias:t" pg:"media,discard_unknown_columns"`

	ID        int     `sql:"id,pk"`
	Thumbnail *string `sql:"thumbnail"`
	Raw       *string `sql:"raw"`
}
