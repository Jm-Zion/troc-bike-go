//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

var Columns = struct {
	Account struct {
		ID, Type string
	}
	BikeOffer struct {
		ID, Size, WheelSize, OfferID, ElectricAssist string

		Offer string
	}
	Category struct {
		ID, Title, Description, ParentID string

		Parent string
	}
	Gear struct {
		ID, Title, Description, Category string

		CategoryRel, CategoryRel1 string
	}
	Location struct {
		ID, City, Address, Zip, Country string
	}
	Login struct {
		ID, Login, Password, UserID string

		User string
	}
	Media struct {
		ID, Thumbnail, Raw string
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
	User struct {
		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

		Account string
	}
}{
	Account: struct {
		ID, Type string
	}{
		ID:   "id",
		Type: "type",
	},
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
	Location: struct {
		ID, City, Address, Zip, Country string
	}{
		ID:      "id",
		City:    "city",
		Address: "address",
		Zip:     "zip",
		Country: "country",
	},
	Login: struct {
		ID, Login, Password, UserID string

		User string
	}{
		ID:       "id",
		Login:    "login",
		Password: "password",
		UserID:   "user_id",

		User: "User",
	},
	Media: struct {
		ID, Thumbnail, Raw string
	}{
		ID:        "id",
		Thumbnail: "thumbnail",
		Raw:       "raw",
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

		Author:    "AuthorRel",
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
	User: struct {
		ID, Firstname, Lastname, MobilePhone, Email, AccountID, UpdatedAt, DeletedAt, CreatedAt string

		Account string
	}{
		ID:          "id",
		Firstname:   "firstname",
		Lastname:    "lastname",
		MobilePhone: "mobile_phone",
		Email:       "email",
		AccountID:   "account_id",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		CreatedAt:   "created_at",

		Account: "Account",
	},
}

var Tables = struct {
	Account struct {
		Name, Alias string
	}
	BikeOffer struct {
		Name, Alias string
	}
	Category struct {
		Name, Alias string
	}
	Gear struct {
		Name, Alias string
	}
	Location struct {
		Name, Alias string
	}
	Login struct {
		Name, Alias string
	}
	Media struct {
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
	User struct {
		Name, Alias string
	}
}{
	Account: struct {
		Name, Alias string
	}{
		Name:  "account",
		Alias: "t",
	},
	BikeOffer: struct {
		Name, Alias string
	}{
		Name:  "bike_offer",
		Alias: "t",
	},
	Category: struct {
		Name, Alias string
	}{
		Name:  "category",
		Alias: "t",
	},
	Gear: struct {
		Name, Alias string
	}{
		Name:  "gear",
		Alias: "t",
	},
	Location: struct {
		Name, Alias string
	}{
		Name:  "location",
		Alias: "t",
	},
	Login: struct {
		Name, Alias string
	}{
		Name:  "login",
		Alias: "t",
	},
	Media: struct {
		Name, Alias string
	}{
		Name:  "media",
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
	User: struct {
		Name, Alias string
	}{
		Name:  "user",
		Alias: "t",
	},
}

