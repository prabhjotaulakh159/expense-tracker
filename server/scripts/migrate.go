package main

import (
	"github.com/prabhjotaulakh159/expense-tracker/models"
	"github.com/prabhjotaulakh159/expense-tracker/db"
	"log"
)

func main() {
	db, err := db.NewGormDB()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}
	defer func() {
		sqlDb, err := db.DB()
		if err != nil {
			log.Fatalf("closing database connection: %v", err)
		}
		if err := sqlDb.Close(); err != nil {
			log.Fatalf("closing database connection: %v", err)
		}
		log.Println("database connection closed")
	}()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	log.Println("migration successful")
}