package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"rohidevs.engineer/mailTrack/Controller"
	"rohidevs.engineer/mailTrack/Utlis"
)

func main() {
	_ = godotenv.Load()
	echoRouter := echo.New()
	echoRouter.Use(middleware.Logger())
	echoRouter.Use(middleware.Recover())
	db, err := DBInit()
	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to database: %s", err))
	}
	err = Utlis.Migrate(db)
	if err != nil {
		slog.Error(fmt.Sprintf("Error migrating database: %s", err))
	}

	//config := echojwt.Config{
	//	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	//		return new(Authentication.Auth)
	//	},
	//	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	//}

	authRouter := echoRouter.Group("/auth")
	//authRouter.Use(echojwt.WithConfig(config))
	Controller.AuthController(authRouter, db)
	echoRouter.Logger.Fatal(echoRouter.Start(":8080"))
}
