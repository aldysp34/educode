package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

type learning struct {
	Title string `json:"learning_title"`
	Desc  string `json:"learning_desc"`
}

func GetLearningByClass(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id := c.QueryParam("class_id")

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

func GetLearning(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	var learning models.Learning
	id := c.QueryParam("learning_id")

	class_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	uint_class_id := uint(class_id)

	if result := database.Db.Where(&models.Learning{LearningID: uint_class_id}).First(&learning); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieve",
		"payload": learning,
	})
}

func CreateNewLearning(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	var newLearning models.Learning

	id := c.QueryParam("class_id")
	class := new(learning)
	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}

	class_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	uint_class_id := uint(class_id)

	newLearning.ClassID = uint_class_id
	newLearning.Title = class.Title
	newLearning.Desc = class.Desc

	if result := database.Db.Model(&models.Learning{}).Create(&newLearning); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Create New Learning Successfully",
		"payload": newLearning,
	})
}

func UpdateLearning(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	class := new(learning)
	id := c.QueryParam("learning_id")
	var updatedLearning models.Learning
	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}
	learning_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_learning_id := uint(learning_id)

	// check database based on id
	if result := database.Db.Where(&models.Learning{LearningID: uint_learning_id}).First(&updatedLearning); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	updatedLearning.Desc = class.Desc
	updatedLearning.Title = class.Title

	if err := database.Db.Save(&updatedLearning); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Update Learning Successfully",
		"payload": updatedLearning,
	})
}

func DeleteLearning(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	id := c.QueryParam("learning_id")
	learning_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_learning_id := uint(learning_id)
	if result := database.Db.Delete(&models.Learning{}, uint_learning_id); result.Error != nil {
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
