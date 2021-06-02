package auth

import (
	"net/http"
	"os"
	"fmt"
	"time"
	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {

	login := new(Login)

	if err := c.Bind(login); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error while binding values",
			"token": "",
		})
	}

	dbLogin, err := doesLoginExists(login.Login);

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Login does not exists",
			"token": "",
		})
	}

	passwordAreEquals := comparePasswords(dbLogin.Password, []byte(login.Password))

	if passwordAreEquals == false {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "Wrong Credentials",
			"token": "",
		})
	}

	token, err := createJWTToken(login.Login)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n Error: %v \n", err)
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	cookie := &http.Cookie{}
	cookie.Name = "JWTToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token":   token,
	})

}