package models

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Name     string `gorm:"not null"`
	Is_Admin bool
}

type SignInInput struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type SignUpInput struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type UserResponse struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token"`
}
