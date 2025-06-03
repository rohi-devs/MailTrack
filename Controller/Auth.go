package Controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"rohidevs.engineer/mailTrack/Service/Authentication"
	"rohidevs.engineer/mailTrack/Utlis/Middleware"
)

func AuthController(e *echo.Group, db *gorm.DB) {
	e.Use(Middleware.DBMiddleware(db))
	e.POST("/login", Authentication.Login)
	e.POST("/register", Authentication.Register)
}
