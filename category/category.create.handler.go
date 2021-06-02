package category

import (
	"fmt"
	"net/http"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
)


func CreateCategory(c echo.Context) error {

	category := new(Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create category 1, %+v.", err),
		})
	}

	_, err := app.PGMain().Model(category).Insert()

	if err != nil {
		fmt.Errorf("No category found for type %v", category)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error cannot create category 2, %v.", err),
		})
	}
	return c.JSON(http.StatusOK, category)
}