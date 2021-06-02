package location

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)

func DeleteLocation(c echo.Context) error {

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

	_, err = app.PGMain().Model(location).WherePK().Delete()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("No location found for type %v", location),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Location with id : %d deleted successfully.", intLocationID),
	})
}