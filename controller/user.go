package controller

import (
	"net/http"

	"github.com/aldysp34/educode/controller/auth"
	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*auth.JWTClaims)
	user_id := claims.User_id

	var user models.User
	if result := database.Db.Where(&models.User{UserID: uint(user_id)}).First(&user); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  http.StatusNotFound,
			"message": "Invalid ID",
		})
	}

	return c.JSON(http.StatusOK, user)

}
