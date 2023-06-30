package database

import (
	"log"

	"github.com/aldysp34/educode/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open("root:@/educode?parseTime=true"), &gorm.Config{})

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

	return
}
