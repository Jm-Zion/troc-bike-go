package location

import (
	"fmt"
	"net/http"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
)


func CreateLocation(c echo.Context) error {

	location := new(Location)

	if err := c.Bind(location); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create location, %+v.", err),
		})
	}

	locationPoint := fmt.Sprintf("(%s,%s)", *location.Longitude, *location.Latitude)
	location.Point = &locationPoint
	validationErrors, valid := location.Validate()

	if valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create location, %v.", validationErrors),
		})
	}

	_, err := app.PGMain().Model(location).Insert()

	if err != nil {
		fmt.Errorf("No location found for type %v", location)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error cannot create location, %v.", err),
		})
	}

	return c.JSON(http.StatusOK, location)
}