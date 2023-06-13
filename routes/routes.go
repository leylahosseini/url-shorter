package routes

import (
	"leylahosseini/my-url-shorter/controllers"

	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo) {
	e.POST("/shorter", controllers.Shorter)
	e.GET("/:code", controllers.Code)
}
