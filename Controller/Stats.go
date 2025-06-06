package Controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"rohidevs.engineer/mailTrack/Service/Stats"
	"rohidevs.engineer/mailTrack/Utlis/Middleware"
)

func StatsController(e *echo.Group, db *gorm.DB) {
	e.Use(Middleware.DBMiddleware(db))
	e.Use(Middleware.JWTMiddleware())
	e.GET("", Stats.GetCount)
	e.GET("/count", Stats.GetCount)
	e.GET("/count/:id", Stats.GetCountById)
	e.GET("/getCampaigns", Stats.GetCampaign)
}
