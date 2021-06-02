package offer


import (
	"fmt"
	"net/http"
	"time"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/Jm-Zion/troc-bike-go/user"
	"github.com/labstack/echo/v4"
)


func CreateOffer(c echo.Context) error {

	offer := new(Offer)

	if err := c.Bind(offer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create offer, %+v.", err),
		})
	}
	
	user, err := user.ExtractUserFromSession(c)

	dateNow := time.Now()
	defaultEnabled := true

	offer.AuthorID = &user.ID
	offer.CreatedAt = &dateNow
	offer.UpdatedAt = &dateNow
	offer.Enabled = &defaultEnabled
	
	validationErrors, valid := offer.Validate()

	if valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error cannot create offer, %v.", validationErrors),
		})
	}

	_, err = app.PGMain().Model(offer).Insert()

	if err != nil {
		fmt.Errorf("No offer found for type %v", offer)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error cannot create offer, %v.", err),
		})
	}

	return c.JSON(http.StatusOK, offer)
}