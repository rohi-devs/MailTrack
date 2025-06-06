package Controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"rohidevs.engineer/mailTrack/Service/Track"
	"rohidevs.engineer/mailTrack/Utlis/Middleware"
)

func TrackController(e *echo.Group, db *gorm.DB) {
	e.Use(Middleware.DBMiddleware(db))
	e.GET("/:id/:CampId/:UserMail", Track.Track)
}
