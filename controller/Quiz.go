package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type quiz struct {
	Soal    string              `json:"quiz_soal" form:"quiz_soal"`
	Options []models.QuizOption `json:"quiz_options"`
}

func CreateQuiz(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
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

func GetQuiz(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id := c.QueryParam("quiz_id")
	quiz_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}

	uint_quiz_id := uint(quiz_id)

	var quiz models.Quiz
	if result := database.Db.Where(&models.Quiz{QuizID: uint_quiz_id}).Preload(clause.Associations).First(&quiz); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieve",
		"payload": quiz,
	})

}

func UpdateQuiz(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id := c.QueryParam("quiz_id")
	quiz_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}
	quiz_json := new(quiz)

	if err := c.Bind(quiz_json); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}
	uint_quiz_id := uint(quiz_id)

	var quiz models.Quiz
	if result := database.Db.Where(&models.Quiz{QuizID: uint_quiz_id}).First(&quiz); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	quiz.Soal = quiz_json.Soal

	if result := database.Db.Save(&quiz); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Update Quiz SuccessFully",
		"payload": quiz,
	})
}

func DeleteQuiz(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id := c.QueryParam("quiz_id")
	quiz_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}
	uint_quiz_id := uint(quiz_id)

	if result := database.Db.Delete(&models.Quiz{}, uint_quiz_id); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"status":  http.StatusNoContent,
		"message": "Delete Quiz Successfully",
	})
}
