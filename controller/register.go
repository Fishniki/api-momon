package controller

import (
	"net/http"
	"api-momon/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Insert the user into the database
	db := config.GetDB()
	_, err = db.Exec("INSERT INTO users (nama, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}
