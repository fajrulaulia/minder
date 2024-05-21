package helper

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var secretKey = []byte("secret")

func CreateToken(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secretKey)
}

func ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := c.Request().Header.Get("Authorization")

		if payload == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not allowed access this endpoint")
		}
		if len(payload) == 0 || !strings.Contains(payload, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not allowed access this endpoint")
		}

		var tokenClaims jwt.MapClaims

		_, err := jwt.ParseWithClaims(payload[7:], &tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("username", tokenClaims["username"].(string))
		c.Set("email", tokenClaims["email"].(string))
		return next(c)
	}
}
