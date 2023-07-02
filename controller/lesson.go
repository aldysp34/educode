package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type LessonStruct struct {
	Title string `json:"lesson_title"`
	Text  string `json:"lesson_text"`
}
type QuizStruct struct {
}

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

func CreateNewLesson(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	var newLearning models.Lesson
	id := c.Param("learning_id")

	learning_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	uint_learning_id := uint(learning_id)

	lessonStruct := new(LessonStruct)
	if err := c.Bind(lessonStruct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}

	newLearning.LessonID = uint_learning_id
	newLearning.Title = lessonStruct.Title
	newLearning.Text = lessonStruct.Text

	if result := database.Db.Create(&newLearning); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	// Create Quiz and File

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Create New Learning Successfully",
		"payload": newLearning,
	})
}

func UpdateLesson(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	lessonStruct := new(LessonStruct)
	var updatedLearning models.Lesson
	id := c.Param("learning_id")
	learning_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_learning_id := uint(learning_id)

	// check database based on id
	if result := database.Db.Where(&models.Lesson{LessonID: uint_learning_id}).First(&updatedLearning); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	if err := c.Bind(lessonStruct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}

	updatedLearning.Text = lessonStruct.Text
	updatedLearning.Title = lessonStruct.Title

	if err := database.Db.Save(&updatedLearning); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error,
		})
	}

	// Update File and Media

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Update Learning Successfully",
		"payload": updatedLearning,
	})
}

func DeleteLesson(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	id := c.Param("learning_id")
	learning_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_learning_id := uint(learning_id)
	if result := database.Db.Delete(&models.Lesson{}, uint_learning_id); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"status":  http.StatusNoContent,
		"message": "Delete Learning Successfully",
	})
}
