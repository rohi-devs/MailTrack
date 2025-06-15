package Track

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log/slog"
	"rohidevs.engineer/mailTrack/Model"
)

func Track(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	userId, _ := uuid.Parse(id)
	campId := c.Param("CampId")
	userMail := c.Param("UserMail")
	slog.Info(fmt.Sprintf("Tracking user mail: %s for the Campagin %s by ID %s", userMail, campId, id))
	event := &Model.TrackEvent{
		UserMail: userMail,
		CampId:   campId,
		UserID:   userId,
		IpAddr:   c.RealIP(),
	}
	if err := db.Create(event).Error; err != nil {
		slog.Error("Failed to create track event", "error", err)
		return c.JSON(500, echo.Map{
			"message": "Internal server error",
		})
	}
	imgPath := "asset/1x1.png"
	return c.File(imgPath)
}
