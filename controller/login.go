package controller

import (
	"api-momon/config"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)


type User2 struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func LoginUser(c echo.Context) error {
	user := new(User2)

	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Terdapat error")
	}
	

	var storeUser User
	db := config.GetDB()

	err := db.QueryRow("SELECT id, nama, email, password FROM users WHERE email = ?", user.Email).Scan(&storeUser.Id, &storeUser.Name, &storeUser.Email, &storeUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusUnauthorized, "User tidak dikenali")
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "User tidak dikenali")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"nama" : storeUser.Name,
		"email" : storeUser.Email,
		"id" : storeUser.Id,
	})

}