package offer

import (
	"time"
	user "github.com/Jm-Zion/troc-bike-go/user"
	"github.com/Jm-Zion/troc-bike-go/gear"
	"github.com/Jm-Zion/troc-bike-go/category"
	"github.com/Jm-Zion/troc-bike-go/location"
	"github.com/Jm-Zion/troc-bike-go/media"	
)

type OfferDesign struct {
	tableName struct{} `sql:"offer_design,alias:t" pg:"offer_design,discard_unknown_columns"`

	ID         int     `sql:"id,pk"`
	Theme *string `sql:"theme"`
	Color      *string `sql:"color"`
}


type OfferGear struct {
	tableName struct{} `sql:"offer_gear,alias:t" pg:",discard_unknown_columns"`

	ID      int  `sql:"id,pk"`
	OfferID *int `sql:"offer_id"`
	GearID  *int `sql:"gear_id"`

	Gear  *gear.Gear  `pg:"fk:gear_id"`
	Offer *Offer `pg:"fk:offer_id"`
}

type OfferMedia struct {
	tableName struct{} `sql:"offer_media,alias:t" pg:",discard_unknown_columns"`

	ID      int  `sql:"id,pk"`
	OfferID *int `sql:"offer_id"`
	MediaID *int `sql:"media_id"`

	Media *media.Media `pg:"fk:media_id"`
	Offer *Offer  `pg:"fk:offer_id"`
}

type BikeOffer struct {
	tableName struct{} `sql:"bike_offer,alias:t" pg:"bike_offer,discard_unknown_columns"`

	ID             int     `sql:"id,pk" json:"id"`
	Size           *string `sql:"size" json:"size"`
	WheelSize      *string `sql:"wheel_size" json:"wheelSize"`
	OfferID        *int    `sql:"offer_id" json:"offerId"`
	ElectricAssist *bool   `sql:"electric_assist" json:"electric"`

	Offer *Offer `pg:"fk:offer_id" json:"offer"`
}

type Offer struct {
	tableName struct{} `sql:"offer,alias:t" pg:"offer,discard_unknown_columns"`

	ID          int        `sql:"id,pk" json:"id"`
	AuthorID    *int      `sql:"author_id" json:"author_id"`
	Title       *string    `sql:"title" json:"title"`
	Description *string    `sql:"description" json:"description"`
	Price       *int       `sql:"price" json:"price"`
	Negociation *bool      `sql:"negociation" json:"negociation"`
	Category    *int       `sql:"category" json:"category_id"`
	Design      *int       `sql:"design" json:"design_id"`
	UpdatedAt   *time.Time `sql:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `sql:"deleted_at" json:"deleted_at"`
	CreatedAt   *time.Time `sql:"created_at" json:"created_at"`
	LocationID  *int       `sql:"location_id" json:"location_id"`
	Size        *string    `sql:"size" json:"size"`
	Condition   *int       `sql:"condition" json:"condition"`
	Enabled     *bool      `sql:"enabled" json:"enabled"`
	Quantity    *int       `sql:"quantity" json:"quantity"`

	Author    *user.User        `pg:"fk:author_id" json:"author"`
	CategoryRel1 *category.Category    `pg:"fk:category" json:"category"`
	DesignRel    *OfferDesign `pg:"fk:design"`
	Location     *location.Location    `pg:"fk:location_id" json:"location"`
}


var Columns = struct {
	BikeOffer struct {
		ID, Size, WheelSize, OfferID, ElectricAssist string

		Offer string
	}
	Offer struct {
		ID, AuthorID, Title, Description, Price, Negociation, Category, Design, UpdatedAt, DeletedAt, CreatedAt, LocationID, WheelSize, Size, Condition, Enabled, Quantity string

		Author, CategoryRel, CategoryRel1, DesignRel, Location string
	}
	OfferDesign struct {
		ID, Theme, Color string
	}
	OfferGear struct {
		ID, OfferID, GearID string

		Gear, Offer string
	}
	OfferMedia struct {
		ID, OfferID, MediaID string

		Media, Offer string
	}

}{
	BikeOffer: struct {
		ID, Size, WheelSize, OfferID, ElectricAssist string

		Offer string
	}{
		ID:             "id",
		Size:           "size",
		WheelSize:      "wheel_size",
		OfferID:        "offer_id",
		ElectricAssist: "electric_assist",

		Offer: "Offer",
	},
	Offer: struct {
		ID, AuthorID, Title, Description, Price, Negociation, Category, Design, UpdatedAt, DeletedAt, CreatedAt, LocationID, WheelSize, Size, Condition, Enabled, Quantity string

		Author, CategoryRel, CategoryRel1, DesignRel, Location string
	}{
		ID:          "id",
		AuthorID:      "author",
		Title:       "title",
		Description: "description",
		Price:       "price",
		Negociation: "negociation",
		Category:    "category",
		Design:      "design",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		CreatedAt:   "created_at",
		LocationID:  "location_id",
		WheelSize:   "wheel_size",
		Size:        "size",
		Condition:   "condition",
		Enabled:     "enabled",
		Quantity:    "quantity",

		Author:    "Author",
		CategoryRel:  "CategoryRel",
		CategoryRel1: "CategoryRel1",
		DesignRel:    "DesignRel",
		Location:     "Location",
	},
	OfferDesign: struct {
		ID, Theme, Color string
	}{
		ID:         "id",
		Theme: "theme",
		Color:      "color",
	},
	OfferGear: struct {
		ID, OfferID, GearID string

		Gear, Offer string
	}{
		ID:      "id",
		OfferID: "offer_id",
		GearID:  "gear_id",

		Gear:  "Gear",
		Offer: "Offer",
	},
	OfferMedia: struct {
		ID, OfferID, MediaID string

		Media, Offer string
	}{
		ID:      "id",
		OfferID: "offer_id",
		MediaID: "media_id",

		Media: "Media",
		Offer: "Offer",
	},
}

var Tables = struct {
	BikeOffer struct {
		Name, Alias string
	}
	Offer struct {
		Name, Alias string
	}
	OfferDesign struct {
		Name, Alias string
	}
	OfferGear struct {
		Name, Alias string
	}
	OfferMedia struct {
		Name, Alias string
	}
}{
	BikeOffer: struct {
		Name, Alias string
	}{
		Name:  "bike_offer",
		Alias: "t",
	},
	Offer: struct {
		Name, Alias string
	}{
		Name:  "offer",
		Alias: "t",
	},
	OfferDesign: struct {
		Name, Alias string
	}{
		Name:  "offer_design",
		Alias: "t",
	},
	OfferGear: struct {
		Name, Alias string
	}{
		Name:  "offer_gear",
		Alias: "t",
	},
	OfferMedia: struct {
		Name, Alias string
	}{
		Name:  "offer_media",
		Alias: "t",
	},
}

