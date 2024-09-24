package DB

import (
	"WhatsAppClone/Helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPG() (*gorm.DB, error) {
	dsn := Helpers.GetEnv("dsnPG")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
