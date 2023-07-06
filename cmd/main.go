package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aldysp34/educode/controller"
	"github.com/aldysp34/educode/controller/auth"
	"github.com/aldysp34/educode/database"
	"github.com/gofiber/fiber/middleware"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type config struct {
	configure echojwt.Config
}

func main() {
	// Initialize Database
	database.Init()

	// Initialize echo framework
	e := echo.New()

	/* Middleware Logger and Recover */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	/* End Middleware */

	// initialize jwt middleware configure
	config := config{create_config()}

	/* Routes */
	api := e.Group("/api")
	authentication := api.Group("/auth")
	authentication.POST("/login", auth.Login)
	authentication.POST("/register", auth.Register)

	user := api.Group("/user")
	{
		user.Use(echojwt.WithConfig(config.configure))

		user.GET("/", controller.GetUser)
	}

	class := api.Group("/class")
	{
		class.Use(echojwt.WithConfig(config.configure))

		class.GET("", controller.GetClasses)
		class.GET("/", controller.GetClass)
		class.POST("/new-class", controller.CreateNewClass)
		class.PUT("/update-class", controller.UpdateClass)
		class.DELETE("/delete-class", controller.DeleteClass)
	}

	learning := class.Group("/learning")
	{
		learning.Use(echojwt.WithConfig(config.configure))

		learning.GET("", controller.GetLearningByClass)
		learning.GET("/", controller.GetLearning)
		learning.POST("/new-learning", controller.CreateNewLearning)
		learning.PUT("/update-learning", controller.UpdateLearning)
	}
	lesson := learning.Group("/lesson")
	{
		lesson.Use(echojwt.WithConfig(config.configure))

		lesson.GET("", controller.GetLessonsByLearningID)
		lesson.GET("/", controller.GetLesson)
		lesson.POST("/new-lesson", controller.CreateNewLesson)
		lesson.PUT("/update-lesson", controller.UpdateLesson)
		lesson.DELETE("/delete-lesson", controller.DeleteLesson)
	}

	quiz := lesson.Group("/quiz")
	{
		quiz.Use(echojwt.WithConfig(config.configure))

		quiz.POST("/new-quiz", controller.CreateQuiz)
	}

	files := lesson.Group("/files")
	{
		files.Use(echojwt.WithConfig(config.configure))

		files.GET("/asset", controller.GetFile)
		files.POST("/new-files", controller.CreateFiles)
	}
	/* End Routes */

	/* Write Route List and export as json */
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	os.WriteFile("routes.json", data, 0644)
	/* End Routes List */

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))

}

// Function to create jwt middleware
func create_config() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}

	return config
}
