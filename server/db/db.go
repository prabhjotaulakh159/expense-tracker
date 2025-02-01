package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"fmt"
)

func NewGormDB() (*gorm.DB, error) {
	host := os.Getenv("pg_host")
	port := os.Getenv("pg_port")
	user := os.Getenv("pg_username")
	password := os.Getenv("pg_password")
	dbname := os.Getenv("pg_database_name")
	dsnf := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn := fmt.Sprintf(dsnf, host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}