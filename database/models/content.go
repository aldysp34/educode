package models

type LearningClass struct {
	LearningClassID uint       `gorm:"primaryKey" json:"class_id"`
	Title           string     `json:"class_title"`
	Desc            string     `json:"class_desc"`
	Learnings       []Learning `gorm:"foreignKey:ClassID;references:LearningClassID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
}

type Learning struct {
	LearningID uint   `gorm:"primaryKey" json:"learning_id"`
	Title      string `json:"learning_title"`
	Desc       string `json:"learning_desc"`
	ClassID    uint
	Lessons    []Lesson `gorm:"foreignKey:LearnID;references:LearningID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
}

type Lesson struct {
	LessonID  uint `gorm:"primaryKey" json:"lesson_id"`
	LearnID   uint
	Is_active bool
	Title     string        `json:"lesson_title"`
	Text      string        `json:"lesson_text"`
	Snippet   string        `json:"lesson_snippet"`
	Notes     string        `json:"lesson_notes"`
	Files     []FileContent `gorm:"foreignKey:LessID;references:LessonID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	Quizzies  []Quiz        `gorm:"foreignKey:LessID;references:LessonID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
}

type FileContent struct {
	ID       uint   `gorm:"primaryKey" json:"file_id"`
	Filetype string `json:"filetype"`
	Filename string `json:"filename"`
	Filesize string `json:"filesize"`
	Filepath string `json:"filepath"`
	LessID   uint
}

type Quiz struct {
	QuizID  uint   `gorm:"primaryKey" json:"quiz_id"`
	Soal    string `json:"quiz_soal"`
	LessID  uint
	Options []QuizOption `gorm:"foreignKey:KuisID;references:QuizID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
}

type QuizOption struct {
	QuizOptionID uint   `gorm:"primaryKey" json:"quiz_option_id"`
	Desc         string `json:"quiz_option_desc"`
	Is_true      bool   `json:"is_true"`
	KuisID       uint
}
