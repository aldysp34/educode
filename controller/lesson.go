package controller

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type LessonJSON struct {
	Title   string `json:"lesson_title" form:"lesson_title"`
	Text    string `json:"lesson_text" form:"lesson_text"`
	Snippet string `json:"lesson_snippet" form:"lesson_snippet"`
}

type LessonStruct struct {
	Title string
	Text  string
	Quiz  []models.Quiz
	Media []models.FileContent
}

type QuizStruct struct {
	QuizID  uint
	LessID  uint
	Options []models.QuizOption
	Soal    string
}

type QuizOptions struct {
	QuizOptionID uint
	Desc         string
	Is_true      bool
	KuisID       uint
}
type successCheck struct {
	IsSuccess bool
	Message   string
}

func (lesson *LessonStruct) AddMedia(media models.FileContent) []models.FileContent {
	lesson.Media = append(lesson.Media, media)
	return lesson.Media
}

func (lesson *LessonStruct) AddQuiz(quiz models.Quiz) []models.Quiz {
	lesson.Quiz = append(lesson.Quiz, quiz)

	return lesson.Quiz
}

func (quiz *QuizStruct) AddOptions(option models.QuizOption) []models.QuizOption {
	quiz.Options = append(quiz.Options, option)

	return quiz.Options
}

func GetLessonsByLearningID(c echo.Context) error {
	id := c.QueryParam("learning_id")

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

func GetLesson(c echo.Context) error {
	id := c.QueryParam("lesson_id")

	learning_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}

	uint_learning_id := uint(learning_id)

	var lesson []models.Lesson
	if result := database.Db.Where(&models.Lesson{LessonID: uint_learning_id}).Preload(clause.Associations).Preload("Quizzies." + clause.Associations).First(&lesson); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "data retrieve",
		"payload": lesson,
	})

}

func CreateNewLesson(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	var newLearning models.Lesson
	id := c.QueryParam("learning_id")

	learning_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	uint_learning_id := uint(learning_id)

	lessonJSON := new(LessonJSON)
	if err := c.Bind(lessonJSON); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}
	// fmt.Printf("JSON: %+v\nData type: %s\n", lessonJSON.Quiz, reflect.TypeOf(lessonJSON))
	// fmt.Println("uint learning id: ", uint_learning_id)

	newLearning.LearnID = uint_learning_id
	newLearning.Title = lessonJSON.Title
	newLearning.Text = lessonJSON.Text
	newLearning.Snippet = lessonJSON.Snippet

	if result := database.Db.Create(&newLearning); result.Error != nil {
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

func UpdateLesson(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()

	lessonJSON := new(LessonJSON)
	var updatedLearning models.Lesson
	id := c.QueryParam("lesson_id")
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

	if err := c.Bind(lessonJSON); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Can't bind request JSON",
		})
	}

	updatedLearning.Text = lessonJSON.Text
	updatedLearning.Title = lessonJSON.Title

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

	id := c.QueryParam("lesson_id")
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
