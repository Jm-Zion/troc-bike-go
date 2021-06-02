package category

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)


func GetCategoryById(c echo.Context) error {

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

	err = app.PGMain().Model(cat).WherePK().Select()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("No category found for type %v", cat),
		})
	}
	return c.JSON(http.StatusOK, cat)
}

func GetRootCategories(c echo.Context) error {

	var categories []Category

	err := app.PGMain().Model(&categories).Where("parent_id IS NULL").Select()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Something wrong happened retrieving root categories"),
		})
	}

	return c.JSON(http.StatusOK, categories)
}

func GetCategoriesForParentID(c echo.Context) error {

	strCatID := c.Param("category_id")

	intCatID, err := strconv.Atoi(strCatID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	var categories []Category

	err = app.PGMain().Model(&categories).Where("parent_id = ?",intCatID).Select()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Something wrong happened retrieving root categories"),
		})
	}

	return c.JSON(http.StatusOK, categories)
}