package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func GetGormInstance() (*gorm.DB, error) {
	if instance == nil {
		host := os.Getenv("pg_host")
		port := os.Getenv("pg_port")
		user := os.Getenv("pg_username")
		password := os.Getenv("pg_password")
		dbname := os.Getenv("pg_database_name")
		connStrFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
		connStr := fmt.Sprintf(connStrFormat, host, user, password, dbname, port)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		instance = db
	}
	return instance, nil
}
