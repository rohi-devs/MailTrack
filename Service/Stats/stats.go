package Stats

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"rohidevs.engineer/mailTrack/Model"
	"rohidevs.engineer/mailTrack/Service/Authentication"
)

func GetCount(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	token := c.Get("AUTH").(*jwt.Token)
	claims := token.Claims.(*Authentication.Auth)
	uri := c.Request().RequestURI
	method := c.Request().Method
	var count int64
	db.Model(&Model.TrackEvent{}).Where("user_id = ?", claims.ID).Count(&count)
	slog.Info(
		fmt.Sprintf("Email: %s with Role: %s -> Accessed: %s with method: %s",
			claims.Email,
			claims.Role,
			uri,
			method,
		),
	)
	return c.JSON(http.StatusOK, echo.Map{
		"UserId": claims.ID,
		"Email":  claims.Email,
		"count":  count,
	})
}

func GetCountById(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	token := c.Get("AUTH").(*jwt.Token)
	claims := token.Claims.(*Authentication.Auth)
	uri := c.Request().RequestURI
	method := c.Request().Method
	id := c.Param("id")
	var count int64
	if err := db.Model(&Model.TrackEvent{}).Where("user_id = ? AND camp_id = ?", claims.ID, id).Count(&count).Error; err != nil {
		slog.Error("Failed to get count by ID", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	slog.Info(
		fmt.Sprintf("Email: %s with Role: %s -> Accessed: %s with method: %s",
			claims.Email,
			claims.Role,
			uri,
			method,
		),
	)
	return c.JSON(http.StatusOK, echo.Map{
		"UserId": claims.ID,
		"Email":  claims.Email,
		"count":  count,
	})
}

func GetCampaign(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	token := c.Get("AUTH").(*jwt.Token)
	claims := token.Claims.(*Authentication.Auth)
	uri := c.Request().RequestURI
	method := c.Request().Method
	var campaign []string
	if err := db.Model(&Model.TrackEvent{}).Distinct("camp_id").Where("user_id = ?", claims.ID).Pluck("camp_id", &campaign).Error; err != nil {
		slog.Error("Failed to get campaigns", "error", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	slog.Info(
		fmt.Sprintf("Email: %s with Role: %s -> Accessed: %s with method: %s",
			claims.Email,
			claims.Role,
			uri,
			method,
		),
	)
	return c.JSON(http.StatusOK, echo.Map{
		"UserId":    claims.ID,
		"Email":     claims.Email,
		"campaigns": campaign,
	})
}
