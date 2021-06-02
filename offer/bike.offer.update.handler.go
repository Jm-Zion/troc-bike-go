package offer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
)

func CreateBikeOffer(c echo.Context) error {

	strOfferID := c.Param("bike_offer_id")

	intOfferID, err := strconv.Atoi(strOfferID)

	bikeOffer := new(BikeOffer)

	if err := c.Bind(bikeOffer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create bike offer, %+v.", err),
		})
	}

	if err, valid := bikeOffer.Validate(); valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	bikeOffer.OfferID = &intOfferID

	validationErrors, valid := bikeOffer.Validate()

	if valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create bike offer, %v.", validationErrors),
		})
	}

	_, err = app.PGMain().Model(bikeOffer).Insert()

	if err != nil {
		fmt.Errorf("No bike offer found for type %v", bikeOffer)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error cannot create bike offer, %v.", err),
		})
	}

	return c.JSON(http.StatusOK, bikeOffer)
}
