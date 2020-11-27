package models

import "gorm.io/gorm"

type RateRepository struct {
	Db *gorm.DB
}

func NewRateRepository(db *gorm.DB) *RateRepository {
	return &RateRepository{
		Db: db,
	}
}

func (rateRepository *RateRepository) AddRate(rate Rate) (int, error) {
	rateRepository.Db.Save(&rate)

	return rate.ID, nil
}
