package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func RunDb() (*gorm.DB, error) {

	dbrUri := "host=localhost port=5432 user=postgres password=******* dbname=diary sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbrUri), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
