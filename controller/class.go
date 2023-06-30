package controller

import (
	"net/http"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

func GetClasses(c echo.Context) error {
	var class []models.LearningClass

	if result := database.Db.Find(&class); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieves",
		"payload": class,
	})
}
