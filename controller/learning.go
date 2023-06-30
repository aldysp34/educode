package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

func GetLearningByClass(c echo.Context) error {
	id := c.Param("class_id")

	class_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	uint_class_id := uint(class_id)

	var learnings []models.Learning
	if result := database.Db.Where(&models.Learning{ClassID: uint_class_id}).Find(&learnings); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieves",
		"payload": learnings,
	})

}
