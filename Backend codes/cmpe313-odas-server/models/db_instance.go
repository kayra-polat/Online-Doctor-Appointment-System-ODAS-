package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func connectDatabase() {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 user=postgres password=Rp5Gy8p3YP dbname=odasDB port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot establish database connection: %s", err.Error())
	}

	db = conn
}

func GetDatabaseInstance() *gorm.DB {
	if db == nil {
		connectDatabase()
	}
	return db
}
