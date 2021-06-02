package user

import (
	"net/http"
	"strconv"
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {

	strUserID := c.Param("user_id")

	user := new(User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error cannot bind user.",
		})
	}

	if err, valid := user.Validate(); valid != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	intUserID, err := strconv.Atoi(strUserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error converting string to integer.",
		})
	}


	userToUpdate := &User{ ID: intUserID}

	err = app.PGMain().Model(userToUpdate).WherePK().Select()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	if user.Firstname != nil {
		userToUpdate.Firstname = user.Firstname
	}
	if user.Lastname != nil {
		userToUpdate.Lastname = user.Lastname
	}
	if user.Username != nil {
		userToUpdate.Username = user.Username
	}
	if user.MobilePhone != nil {
		userToUpdate.MobilePhone = user.MobilePhone
	}
	if user.Email != nil {
		userToUpdate.Email = user.Email
	}
	if user.AccountID != nil {
		userToUpdate.AccountID = user.AccountID
	}

	_, err = app.PGMain().Model(userToUpdate).WherePK().Update()
	
	
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("%v", err),
		})
	}

	return c.JSON(http.StatusOK, userToUpdate)
}