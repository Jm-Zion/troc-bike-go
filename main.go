package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	router "github.com/Jm-Zion/troc-bike-go/router"
	"github.com/Jm-Zion/troc-bike-go/auth"
	"github.com/Jm-Zion/troc-bike-go/user"
	"github.com/Jm-Zion/troc-bike-go/media"
	"github.com/Jm-Zion/troc-bike-go/offer"
	"github.com/Jm-Zion/troc-bike-go/location"
	"github.com/Jm-Zion/troc-bike-go/category"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type JwtClaims struct {
	Name string `json:name`
	jwt.StandardClaims
}

func createJWTToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Store MySeret into a secret place
	token, err := rawToken.SignedString([]byte("MySecret"))
	if err != nil {
		return "", err

	}
	return token, nil
}

func connectToDb(c echo.Context) error {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to reach database",
			"env":     os.Getenv("DATABASE_URL"),
		})
	}
	db := pg.Connect(opt)
	if err := db.Ping(c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "DB not running",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "connected to database",
	})
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	if username == "jack" && password == "1234" {
		token, err := createJWTToken()
		if err != nil {
			log.Println("Error Creating JWT token")
			return c.String(http.StatusInternalServerError, "Something went wrong...")
		}
		cookie := &http.Cookie{}
		cookie.Name = "JWTToken"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Logged in",
			"token":   token,
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": "Wrong credentials",
		"token":   "",
	})
}

func mainJwt(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the top jwt secret")
}

func yallo(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&cat)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't parse json"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "200", "name": cat.Name})
}

func main() {
	godotenv.Load()
	r := router.Init()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1Protected := v1.Group("/protected")
	v1Protected.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:   "cookie:JWTToken",
	}))

	// -- AUTH --
	v1.POST("/sign-up", auth.SignUp)
	v1.POST("/sign-in", auth.SignIn)

	// -- USER --
	v1Protected.GET("/user/me", user.GetUserMe)
	v1Protected.PUT("/user/:user_id", user.UpdateUser)

	// -- CATEGORY --
	v1Protected.GET("/categories/:category_id", category.GetCategoryById)
	v1Protected.DELETE("/categories/:category_id", category.DeleteCategoryById)
	v1Protected.POST("/categories", category.CreateCategory)
	v1Protected.GET("/categories/roots", category.GetRootCategories)
	v1Protected.GET("/categories/:category_id/childs", category.GetCategoriesForParentID)

	// -- LOCATION --
	v1Protected.GET("/location/:location_id", location.GetLocationById)
	v1Protected.POST("/location", location.CreateLocation)
	v1Protected.DELETE("/location/:location_id", location.DeleteLocation)

	// -- MEDIA --
	v1Protected.POST("/media", media.CreateMedia)
	v1Protected.GET("/media/:media_id", media.GetMediaById)
	v1Protected.DELETE("/media/:media_id", media.DeleteMedia)

	// -- MEDIA --
	v1Protected.POST("/offers", offer.CreateOffer)
	v1Protected.GET("/offers",offer.GetAllOffers)
	v1Protected.PUT("/offers/:offer_id", offer.UpdateOffer)
	v1Protected.POST("/offers/:offer_id/bike", offer.CreateBikeOffer)
	v1Protected.GET("/offers/designs", offer.GetAllOfferDesigns)
	v1Protected.GET("/offers/search", offer.SearchOffers)
	v1Protected.GET("/offers/bike/search",offer.SearchBikeOffers)
	v1Protected.GET("/offers/designs/:offer_design_id", offer.GetOfferDesignById)
		
	v1.GET("/main", mainJwt)
	v1.GET("/db", connectToDb)
	v1Protected.POST("/cats", yallo)
	api.GET("/login", login)
	r.Start(":8000")
	r.Logger.Fatal(gracehttp.Serve(r.Server))
}
