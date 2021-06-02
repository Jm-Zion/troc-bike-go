package user

import (
	"net/http"
	"os"
	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/labstack/echo/v4"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Name string `json:identifier`
	jwt.StandardClaims
}


func ExtractUserFromSession(c echo.Context)(*User, error){
	// Retrieve token from request
	cookie, err := c.Cookie("JWTToken")

	if err != nil {
		return nil, err
	}

	claims := JwtClaims{}
	
	_, err = jwt.ParseWithClaims(cookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	user := &User{
		MobilePhone: &claims.Name,
	}

	err = app.PGMain().Model(user).Where("mobile_phone = ?", claims.Name).Select()

	return user, nil
}


func GetUserMe(c echo.Context) error {

	user, err := ExtractUserFromSession(c)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found.",
		})
	}

	account := &Account{}

	err = app.PGMain().Model(account).Where("id = ?", user.AccountID).Select()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Account not found",
		})
	}

	user.Account = account

	return c.JSON(http.StatusOK, user)

}