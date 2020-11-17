package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
NewDB create a database instance and return to the caller
*/
func NewDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
