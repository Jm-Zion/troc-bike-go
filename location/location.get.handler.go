package location

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)


func GetLocationById(c echo.Context) error {

	strLocationID := c.Param("location_id")

	intLocationID, err := strconv.Atoi(strLocationID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	location := &Location{
		ID: intLocationID,
	}

	err = app.PGMain().Model(location).WherePK().Select()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("No location found for type %v", location),
		})
	}
	return c.JSON(http.StatusOK, location)
}