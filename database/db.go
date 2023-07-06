package database

import (
	"log"
	"os"

	"github.com/aldysp34/educode/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	dbstring := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("HOST_NAME") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dbstring), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect with DB : ", err)
	}
	Db = db

	// Migrate
	db.AutoMigrate(&models.User{},
		&models.LearningClass{},
		&models.Learning{},
		&models.Lesson{},
		&models.FileContent{},
		&models.Quiz{},
		&models.QuizOption{},
	)

	db.Exec("ALTER TABLE file_contents ADD filepath varchar(128);")
	db.Exec("ALTER TABLE lessons ADD snippet longtext")
	return
}
