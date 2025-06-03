package Utlis

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func DBInit() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	d, rr := db.DB()
	if rr != nil {
		slog.Info("Cannot connect to DB:", rr)
	}
	rr = d.Ping()
	if rr != nil {
		slog.Info("Cannot connect to DB:", rr)
		os.Exit(99)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL")
	return db, err
}
