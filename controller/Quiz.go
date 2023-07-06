package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

type quiz struct {
	Soal    string              `json:"quiz_soal" form:"quiz_soal"`
	Options []models.QuizOption `json:"quiz_options"`
}

func CreateQuiz(c echo.Context) error {
	id := c.QueryParam("lesson_id")

	learning_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}

	uint_learning_id := uint(learning_id)

	quiz := new(quiz)

	if err := c.Bind(quiz); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}

	if result := database.Db.Create(&models.Quiz{LessID: uint_learning_id, Soal: quiz.Soal, Options: quiz.Options}); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  http.StatusCreated,
		"message": "Create Quiz Successfully",
		"payload": quiz,
	})
}
