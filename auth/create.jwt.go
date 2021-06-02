package auth

import (
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Name string `json:identifier`
	jwt.StandardClaims
}

func createJWTToken(identifier string) (string, error) {
	claims := JwtClaims{
		identifier,
		jwt.StandardClaims{
			Id:        "main_user",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Store MySeret into a secret place
	token, err := rawToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err

	}
	return token, nil
}
