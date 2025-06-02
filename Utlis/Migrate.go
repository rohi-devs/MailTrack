package Utlis

import (
	"gorm.io/gorm"
	"rohidevs.engineer/mailTrack/Model"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Model.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
