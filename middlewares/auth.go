package middlewares

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

type JWTCustomClaims struct {
	ID int `Json:"id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey string
}

func (jwtConfig *JWTConfig) int() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)

		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, errors.New("invalid token")
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := GetUser(c)

		isInvalid := userData == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"massage": "invalid token",
			})
		}

		return next(c)
	}
}
