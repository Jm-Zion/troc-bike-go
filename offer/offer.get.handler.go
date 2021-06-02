package offer

import (
	"fmt"
	"net/http"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"github.com/go-pg/pg/v10/orm"
	"strconv"
)

type OffersResults struct {
	Data []Offer `json:"data"`
	Total int `json:"total"`
	NextPage int `json:"nextPage"`
	LastPage int `json:"lastPage"`
}

func GetAllOffers(c echo.Context) error {
	var offers []Offer

	pageParam := c.QueryParam("page")
	itemsPerPageParam := c.QueryParam("itemsPerPage")

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

	count, err := app.PGMain().Model(&offers).
		Order("created_at ASC").
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

	return c.JSON(http.StatusOK, &OffersResults{
		Data: offers,
		Total: count,
		NextPage: nextPage,
		LastPage: lastPage,
	})
}

func SearchBikeOffers(c echo.Context) error {
	var offers []Offer

	query := c.QueryParam("query")
	pageParam := c.QueryParam("page")
	itemsPerPageParam := c.QueryParam("itemsPerPage")

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

	count, err := app.PGMain().Model(&offers).
		Where("size = ?", "M").
		WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.WhereOr("to_tsvector(title) @@ to_tsquery('"+query+":*')").
				WhereOr("to_tsvector(description) @@ to_tsquery('"+query+":*')")
			return q, nil
		}).
		Order("created_at ASC").
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

	return c.JSON(http.StatusOK, &OffersResults{
		Data: offers,
		Total: count,
		NextPage: nextPage,
		LastPage: lastPage,
	})
}