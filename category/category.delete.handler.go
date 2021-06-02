package category

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)

func DeleteCategoryById(c echo.Context) error {

	strCatID := c.Param("category_id")

	intCatID, err := strconv.Atoi(strCatID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	cat := &Category{
		ID: intCatID,
	}

	_, err = app.PGMain().Model(cat).WherePK().Delete()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("No category found for type %v", cat),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Category deleted successfully",
	});
}