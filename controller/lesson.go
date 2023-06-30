package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func GetLessonsByLearningID(c echo.Context) error {
	id := c.Param("learning_id")

	learning_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}

	uint_learning_id := uint(learning_id)

	var lessons []models.Lesson
	if result := database.Db.Where(&models.Lesson{LearnID: uint_learning_id}).Preload(clause.Associations).Preload("Quizzies." + clause.Associations).Find(&lessons); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieves",
		"payload": lessons,
	})
}
