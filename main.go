package main

import (
	"api-momon/config"
	"api-momon/controller"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the database connection
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Koneksi Berhasil")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Define the route
	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)


	// Start the server
	e.Logger.Fatal(e.Start(":1234"))
}
