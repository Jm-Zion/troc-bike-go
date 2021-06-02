package auth

import (
	"net/http"
	"os"
	"fmt"
	"time"
	"regexp"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/Jm-Zion/troc-bike-go/user"
	"github.com/labstack/echo/v4"
)


func SignUp(c echo.Context) error {

	login := new(Login)
	
	if err := c.Bind(login); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error while binding values",
			"token": "",
		})
	}

	re := regexp.MustCompile(`^(?:(?:\+|00)33[\s.-]{0,3}(?:\(0\)[\s.-]{0,3})?|0)[1-9](?:(?:[\s.-]?\d{2}){4}|\d{2}(?:[\s.-]?\d{3}){2})$`)
	
	if re.MatchString(login.Login) != true {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Identifier is not a valid phone number",
			"token":   "",
		})
	}

	// Create FREE account by default
	account, err := user.GetAccountByType(user.FREE_ACCOUNT)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating account: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Something went wrong",
			"token":   "",
		})
	}

	now := time.Now()

	u := &user.User{
		MobilePhone: &login.Login,
		UpdatedAt: &now,
		CreatedAt: &now,
		AccountID: &account.ID,
	}

	_, err = app.PGMain().Model(u).Insert()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Something went wrong",
			"token":   "",
		})
	}

	login.UserID = &u.ID
	login.Password = hashAndSalt([]byte(login.Password))
	existingLogin, _ := doesLoginExists(login.Login);

	if existingLogin != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "User already exists",
			"token":   "",
		})
	}

	_, err = app.PGMain().Model(login).Insert()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not save user: %v \n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Something went wrong",
			"token":   "",
		})
	}

	token, err := createJWTToken(login.Login)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n Error while creating JWT token: %+v \n", err)
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
