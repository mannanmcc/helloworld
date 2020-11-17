package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
NewDB create a database instance and return to the caller
*/
func NewDB(dataSourceName string) (*gorm.DB, error) {
	log.Println("datasource name ", dataSourceName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
