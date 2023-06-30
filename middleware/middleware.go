package middleware

import (
	"errors"
	"os"

	"github.com/aldysp34/educode/controller/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("SECRET_KEY")),
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.JWTClaims)
		isAdmin := claims.Admin

		if isAdmin == false {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func RestrictedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("JWT Token is missing")
		}
		_, cok := token.Claims.(jwt.MapClaims)
		if !cok {
			return errors.New("Failed to cast claims as jwt.MapClaims")
		}

		return next(c)
	}
}
