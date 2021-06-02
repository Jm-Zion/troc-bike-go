package media

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)

func DeleteMedia(c echo.Context) error {

	strMediaID := c.Param("media_id")

	intMediaID, err := strconv.Atoi(strMediaID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	media := &Media{
		ID: intMediaID,
	}

	err = app.PGMain().Model(media).WherePK().Select()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("No Media found for type %v", media),
		})
	}

	_, err = app.PGMain().Model(media).WherePK().Delete()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("No media found for type %v", media),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Media with id : %d deleted successfully.", intMediaID),
	})
}