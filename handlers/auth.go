package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"structured/models"
	"time"
)

func RegisterUser(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	var user models.User
	if err := db.First(&user, "username = ?", username).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username Already Exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user = models.User{Username: username, Password: string(hashedPassword)}
	db.Create(&user)

	return c.String(http.StatusOK, "User Created!")
}

func LoginAndCreateToken(c echo.Context) error {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	var user models.User
	db := c.Get("db").(*gorm.DB)
	if err := db.First(&user, "username = ?", data.Username).Error; err != nil {
		return c.String(http.StatusBadRequest, "Username Not Found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		return c.String(http.StatusBadRequest, "Invalid Password")
	}

	expirationTime := time.Now().Add(15 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error signing token")
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
