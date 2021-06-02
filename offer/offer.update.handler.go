package offer


import (
	"fmt"
	"net/http"
	"time"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	"strconv"
)


func UpdateOffer(c echo.Context) error {

	strOfferID := c.Param("offer_id")

	offer := new(Offer)

	if err := c.Bind(offer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error cannot bind offer.",
		})
	}

	if err, valid := offer.Validate(); valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	intOfferID, err := strconv.Atoi(strOfferID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}


	offerToUpdate := &Offer{ ID: intOfferID}

	err = app.PGMain().Model(offerToUpdate).WherePK().Select()

	dateNow := time.Now()

	offer.UpdatedAt = &dateNow

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	if offer.Enabled != nil {
		offerToUpdate.Enabled = offer.Enabled
	}
	if offer.Title != nil {
		offerToUpdate.Title = offer.Title
	}
	if offer.Description != nil {
		offerToUpdate.Description = offer.Description
	}
	if offer.Price != nil {
		offerToUpdate.Price = offer.Price
	}
	if offer.Negociation != nil {
		offerToUpdate.Negociation = offer.Negociation
	}
	if offer.Category != nil {
		offerToUpdate.Category = offer.Category
	}
	if offer.Design != nil {
		offerToUpdate.Design = offer.Design
	}
	if offer.Size != nil {
		offerToUpdate.Size = offer.Size
	}
	if offer.Condition != nil {
		offerToUpdate.Condition = offer.Condition
	}
	if offer.Quantity != nil {
		offerToUpdate.Quantity = offer.Quantity
	}

	_, err = app.PGMain().Model(offerToUpdate).WherePK().Update()

	return c.JSON(http.StatusOK, offerToUpdate)
}