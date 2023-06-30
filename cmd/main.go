package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aldysp34/educode/controller"
	"github.com/aldysp34/educode/controller/auth"
	"github.com/aldysp34/educode/database"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type config struct {
	configure echojwt.Config
}

func main() {
	database.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := config{create_config()}

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

		class.GET("/", controller.GetClasses)
	}

	learning := class.Group("/learning")
	{
		learning.Use(echojwt.WithConfig(config.configure))

		learning.GET("/:class_id", controller.GetLearningByClass)
	}
	lesson := learning.Group("/lesson")
	{
		lesson.Use(echojwt.WithConfig(config.configure))

		lesson.GET("/:learning_id", controller.GetLessonsByLearningID)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	os.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":3000"))

}

func create_config() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}

	return config
}
