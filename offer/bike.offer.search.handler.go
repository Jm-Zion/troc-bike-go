package offer

import (
	"fmt"
	"net/http"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"github.com/go-pg/pg/v10/orm"
	"strconv"
)

type SearchBikeOffersResults struct {
	Data []BikeOffer `json:"data"`
	Total int `json:"total"`
	NextPage int `json:"nextPage"`
	LastPage int `json:"lastPage"`
}

func SearchOffers(c echo.Context) error {
	var offers []BikeOffer

	query := c.QueryParam("query")
	pageParam := c.QueryParam("page")
	size := c.QueryParam("size")
	category := c.QueryParam("category")
	condition := c.QueryParam("condition")
	priceMax := c.QueryParam("priceMax")
	priceMin := c.QueryParam("priceMin")
	itemsPerPageParam := c.QueryParam("itemsPerPage")
	wheelSize := c.QueryParam("wheelSize")
	distance := c.QueryParam("distance")
	latitudeReference := c.QueryParam("latitude")
	longitudeReference := c.QueryParam("longitude")

	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Query param cannot be null.",
		})
	}

	var err error

	// Initial page value
	page := 0
	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
	}

	// Initial itemsPerPage value
	itemsPerPage := 20
	if itemsPerPageParam != "" {
		itemsPerPage, err = strconv.Atoi(itemsPerPageParam)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	//count, err
	sqlQuery := app.PGMain().Model(&offers).
			Relation("Offer").
			Relation("Offer.Location").
			Relation("Offer.Author").
			WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.WhereOr("to_tsvector(offer.title) @@ to_tsquery('"+query+":*')").
				WhereOr("to_tsvector(offer.description) @@ to_tsquery('"+query+":*')")
			return q, nil
		})

	if size != "" {
		sqlQuery = sqlQuery.Where("bike_offer.size = ?", "M")
	}

	if wheelSize != "" {
		sqlQuery = sqlQuery.Where("bike_offer.wheel_size = ?", wheelSize)
	}

	if category != "" {
		sqlQuery = sqlQuery.Where("offer.category = ?", category)
	}

	if condition != "" {
		sqlQuery = sqlQuery.Where("offer.condition = ?", condition)
	}

	if priceMax != "" && priceMin != "" {
		sqlQuery = sqlQuery.Where("offer.price >= ?", priceMin).Where("offer.price =< ?", priceMax)
	}
	if priceMax != "" && priceMin == "" {
		sqlQuery = sqlQuery.Where("offer.price <= ?", priceMax)
	}
	if priceMax == "" && priceMin != "" {
		sqlQuery = sqlQuery.Where("offer.price >= ?", priceMin)
	}

	if distance != "" && latitudeReference != "" && longitudeReference != "" {
		//latitude, err := strconv.ParseFloat(latitudeReference, 32);

		sqlQuery  = sqlQuery.Where("ST_Distance_Sphere(geometry(location.point), geometry(ST_MakePoint("+longitudeReference+","+latitudeReference+"))) <= "+distance)

	}

	//sqlQuery = sqlQuery.Join("INNER JOIN offer ON bike_offer.offer_id = offer.id")

	if distance != "" && latitudeReference != "" && longitudeReference != "" {
		sqlQuery = sqlQuery.Join("INNER JOIN location ON location.id = offer.location_id")
	}

	count, err := sqlQuery.Order("offer.created_at ASC").
		Limit(itemsPerPage).
		Offset(page*itemsPerPage).
		SelectAndCount()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Something wrong happened retrieving offers : %+v",err),
		})
	}

	nextPage := page+1
	if nextPage >= (count/itemsPerPage) {
		nextPage = 0
	}

	lastPage := (count/itemsPerPage) - 1
	if lastPage < 0 {
		lastPage = 0
	}

	return c.JSON(http.StatusOK, &SearchBikeOffersResults{
		Data: offers,
		Total: count,
		NextPage: nextPage,
		LastPage: lastPage,
	})
}