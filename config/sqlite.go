package config

import (
	"os"

	"github.com/lkgiovani/api_rest_golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	
	logger := GetLogger("sqLite")

	_, err := os.Stat(dbPath)
	
	if os.IsNotExist(err) {
		logger.Info("Creating sqLite database...")
		// Create the database file and directory

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file,  err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()

	}

	// create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}) 
	if err != nil {
		logger.ErrF("Error initializing sqLite: %v", err)
		return nil, err
	}

	// Migrate teh schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrF("sqLte automigrate error: %v", err)

		return nil, err
	}

	// return the db


	return db, nil

}