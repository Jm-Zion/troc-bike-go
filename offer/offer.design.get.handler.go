package offer

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
)


func GetOfferDesignById(c echo.Context) error {

	strOfferID := c.Param("offer_design_id")

	intOfferID, err := strconv.Atoi(strOfferID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}

	offer := &OfferDesign{
		ID: intOfferID,
	}

	err = app.PGMain().Model(offer).WherePK().Select()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("No OfferDesign found for type %v", offer),
		})
	}

	return c.JSON(http.StatusOK, offer)
}

func GetAllOfferDesigns(c echo.Context) error {

	var offersDesign []OfferDesign

	_, err := app.PGMain().Model(&offersDesign).SelectAndCount()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Something wrong happened retrieving root categories"),
		})
	}

	return c.JSON(http.StatusOK, offersDesign)
}