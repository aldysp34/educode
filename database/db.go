package database

import (
	"log"

	"github.com/aldysp34/educode/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(connection string) {
	connection = "root:@/educode?parseTime=true"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

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
	db.Exec("ALTER TABLE lessons ADD notes longtext")
	return
}
