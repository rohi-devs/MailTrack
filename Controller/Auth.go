package Controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"rohidevs.engineer/mailTrack/Service/Authentication"
	"rohidevs.engineer/mailTrack/Utlis"
)

func AuthController(e *echo.Group, db *gorm.DB) {
	e.Use(Utlis.DBMiddleware(db))
	e.POST("/login", Authentication.Login)
	e.POST("/register", Authentication.Register)
}
