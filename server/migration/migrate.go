package main

import (
	"log"

	"github.com/prabhjotaulakh159/expense-tracker/db"
	"github.com/prabhjotaulakh159/expense-tracker/models"
)

func main() {
	_db, err := db.GetGormInstance()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// need sql object to close connection
	sqlDb, err := _db.DB()
	if err != nil {
		log.Fatalf("Error getting SQL object: %v", err)
	}
	defer sqlDb.Close()

	if err := _db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed in making migration: %v", err)
	}

	log.Println("Migration complete")
}
