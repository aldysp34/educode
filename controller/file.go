package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/labstack/echo/v4"
)

type Files struct {
	Media []models.FileContent
}

func (f *Files) AddMedia(media models.FileContent) []models.FileContent {
	f.Media = append(f.Media, media)

	return f.Media
}
func CreateFiles(c echo.Context) error {
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

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
	}
	files := form.File["files"]
	if len(files) == 0 {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"status":  http.StatusUnprocessableEntity,
			"message": "No File Uploaded",
		})
	}

	var isSuccess successCheck
	isSuccess.IsSuccess = true

	var fileStruct Files

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			isSuccess.IsSuccess = false
			isSuccess.Message = err.Error()
		} else {
			var media models.FileContent
			fileByte, _ := ioutil.ReadAll(src)
			fileType := http.DetectContentType(fileByte)
			splitExt := strings.Split(fileType, "/")
			extName := splitExt[1]

			fileName := strconv.FormatInt(time.Now().Unix(), 10) + "." + extName
			filePath := "uploads/" + fileName

			err = ioutil.WriteFile(filePath, fileByte, 0777)
			if err != nil {
				isSuccess.IsSuccess = false
				isSuccess.Message = err.Error()
			} else {
				fileSize := file.Size

				media.Filename = fileName
				media.Filesize = strconv.FormatInt(fileSize, 10)
				media.Filetype = fileType
				media.LessID = uint_learning_id
				media.Filepath = filePath

				fileStruct.AddMedia(media)
			}
		}

		if !isSuccess.IsSuccess {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"status":  http.StatusInternalServerError,
				"message": isSuccess.Message,
			})
		}
		defer src.Close()
	}
	for _, x := range fileStruct.Media {
		if result := database.Db.Model(&models.FileContent{}).Create(&x); result.Error != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"status":  http.StatusInternalServerError,
				"message": result.Error,
			})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  http.StatusCreated,
		"message": fileStruct.Media,
	})
}

func GetFile(c echo.Context) error {
	fileLocate := c.QueryParam("path")

	return c.File(fileLocate)
}

func DeleteFile(c echo.Context) error {
	path := c.QueryParam("path")

	if result := database.Db.Delete(&models.FileContent{Filepath: path}); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error Removing file from path ", path)
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"status":  http.StatusNoContent,
		"message": "Delete File Successfully",
	})
}
