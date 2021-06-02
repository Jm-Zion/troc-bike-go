package location


type Location struct {
	tableName struct{} `sql:"location,alias:t" pg:"location,discard_unknown_columns"`

	ID      int     `sql:"id,pk" json:"id"`
	City    *string `sql:"city" json:"city"`
	Address *string `sql:"address" json:"address"`
	Zip     *string `sql:"zip" json:"zip"`
	Point *string `sql:"point" json:"point"`
	Country *string `sql:"country" json:"country"`
	Latitude *string `sql:"latitude" json:"latitude"`
	Longitude *string `sql:"longitude" json:"longitude"`
}

var Columns = struct {
	Location struct {
		ID, City, Address, Zip, Latitude, Longitude, Point, Country string
	}
}{
	Location: struct {
		ID, City, Address, Zip, Latitude, Longitude, Point, Country string
	}{
		ID:      "id",
		City:    "city",
		Address: "address",
		Zip:     "zip",
		Country: "country",
		Point: "point",
		Latitude: "latitude",
		Longitude: "longitude",
	},
}

var Tables = struct {
	Location struct {
		Name, Alias string
	}
}{
	Location: struct {
		Name, Alias string
	}{
		Name:  "location",
		Alias: "t",
	},
}

