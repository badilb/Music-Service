package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"music-service-v2.0.0/internal/config"
	"music-service-v2.0.0/internal/models"
)

var db *gorm.DB

func initializeDB(database config.Database) error {
	dbConnString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		database.User,
		database.Password,
		database.Host,
		database.Port,
		database.Name,
		database.Sslmode,
	)
	fmt.Println(dbConnString)
	var err error
	db, err = gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Playlist{}, &models.Tracks{})
	if err != nil {
		return err
	}

	return nil
}

func GetDBInstance(database config.Database) (*gorm.DB, error) {
	db = nil
	var errGetDB error
	if db == nil {
		if err := initializeDB(database); err != nil {
			errGetDB = err
		}
	}
	return db, errGetDB
}
