package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/aldysp34/educode/database"
	"github.com/aldysp34/educode/database/models"
	"github.com/aldysp34/educode/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	Name    string `json:"name"`
	User_id int    `json:"user_id"`
	Admin   bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func Register(c echo.Context) error {
	var user models.User

	u := new(models.SignUpInput)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request, Can't bind Request")
	}
	if result := database.Db.Where(&models.User{Username: u.Username}).First(&user); result.Error == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Username is already exist",
		})
	}

	user.Name = u.Name
	user.Username = u.Username
	user.Password = utils.HashPassword(u.Password)
	user.Is_Admin = false

	database.Db.Create(&user)

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  http.StatusCreated,
		"message": "Register Successfully",
		"payload": user,
	})
}

func Login(c echo.Context) error {
	var user models.User

	u := new(models.SignInInput)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request, Can't bind Request")
	}
	if result := database.Db.Where(&models.User{Username: u.Username}).First(&user); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  http.StatusNotFound,
			"message": "Username Not Found",
		})
	}

	matchPassword := utils.CheckPasswordHash(u.Password, user.Password)
	if !matchPassword {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid Password",
		})
	}

	// Set custom claims
	claims := &JWTClaims{
		user.Name,
		int(user.UserID),
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	if user.Is_Admin == true {
		claims.Admin = true
	}

	// Create Token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encode token
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	userResponse := models.UserResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    t,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Login Successfully",
		"payload": userResponse,
	})
}
