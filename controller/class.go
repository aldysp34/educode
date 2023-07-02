package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

type class struct {
	Title string `json:"class_title"`
	Desc  string `json:"class_desc"`
}

func GetClasses(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
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

func GetClass(c echo.Context) error {
	Lock.Lock()
	Lock.Unlock()

	id := c.QueryParam("class_id")
	var class models.LearningClass

	class_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	uint_class_id := uint(class_id)

	if result := database.Db.Where(&models.LearningClass{LearningClassID: uint_class_id}).First(&class); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieve",
		"payload": class,
	})
}

func CreateNewClass(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	var newClass models.LearningClass
	class := new(class)
	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}

	newClass.Title = class.Title
	newClass.Desc = class.Desc

	if err := database.Db.Model(&models.LearningClass{}).Create(&newClass); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  http.StatusCreated,
		"message": "Create Learning Class Successfully",
		"payload": newClass,
	})
}

func UpdateClass(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	id := c.QueryParam("class_id")
	class := new(class)
	var updatedClass models.LearningClass

	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}

	class_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_class_id := uint(class_id)

	// check database based on id
	if result := database.Db.Where(&models.LearningClass{LearningClassID: uint_class_id}).First(&updatedClass); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	updatedClass.Desc = class.Desc
	updatedClass.Title = class.Title

	if err := database.Db.Save(updatedClass); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Update Learning Class Successfully",
		"payload": updatedClass,
	})
}

func DeleteClass(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	id := c.QueryParam("class_id")
	class := new(class)
	if err := c.Bind(class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}
	class_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't Convert id")
	}

	uint_class_id := uint(class_id)
	if result := database.Db.Delete(&models.LearningClass{}, uint_class_id); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"status":  http.StatusNoContent,
		"message": "Delete Learning Class Successfully",
	})
}
