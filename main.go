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
	db, err := Utlis.DBInit()
	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to database: %s", err))
	}
	err = Utlis.Migrate(db)
	if err != nil {
		slog.Error(fmt.Sprintf("Error migrating database: %s", err))
	}
	authRouter := echoRouter.Group("/auth")
	statRouter := echoRouter.Group("/stats")
	trackRouter := echoRouter.Group("/track")
	fmt.Println("Starting server on port 8080")
	Controller.AuthController(authRouter, db)
	Controller.StatsController(statRouter, db)
	Controller.TrackController(trackRouter, db)
	echoRouter.Logger.Fatal(echoRouter.Start(":8080"))
}
