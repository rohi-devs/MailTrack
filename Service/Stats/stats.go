package Stats

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"rohidevs.engineer/mailTrack/Service/Authentication"
)

func GetStats(c echo.Context) error {
	token := c.Get("AUTH").(*jwt.Token)
	claims := token.Claims.(*Authentication.Auth)
	uri := c.Request().RequestURI
	method := c.Request().Method
	slog.Info(
		fmt.Sprintf("Email: %s with Role: %s -> Accessed: %s with method: %s",
			claims.Email,
			claims.Role,
			uri,
			method,
		),
	)
	return c.JSON(http.StatusOK, echo.Map{
		"totalEmails": "Invoked GetStats",
		"content":     claims,
	})
}
