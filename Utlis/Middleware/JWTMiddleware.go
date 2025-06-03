package Middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
	"rohidevs.engineer/mailTrack/Service/Authentication"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(Authentication.Auth)
		},
		SigningKey: []byte(jwtSecret),
		ContextKey: "AUTH",
	}
	return echojwt.WithConfig(config)
}
