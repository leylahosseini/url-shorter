package main

import (
	// "leylahosseini/my-url-shorter/db"
	// "leylahosseini/my-url-shorter/db/models"

	"leylahosseini/my-url-shorter/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Initialize Echo
	e := echo.New()
	routes.InitRoute(e)
	// Middleware
	e.Use(middleware.Logger())
	//e.Use(middlewareLogging)
	//e.HTTPErrorHandler = errorHandler

	//e.Use(middleware.Recover())

	// Start server
	e.Start(":8080")

}
