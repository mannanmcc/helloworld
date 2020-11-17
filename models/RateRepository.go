package models

import "gorm.io/gorm"

type RateRepositoryInterface interface {
	AddRate(rate Rate) (int, error)
}

type RateRepository struct {
	Db *gorm.DB
}

/*
AddRate - add rate to the database
*/
func (rateRepository *RateRepository) AddRate(rate Rate) (int, error) {
	rateRepository.Db.Save(&rate)

	return rate.ID, nil
}
