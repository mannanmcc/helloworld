package models

import (
	"fmt"

	"github.com/mannanmcc/helloworld/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Dbname, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
