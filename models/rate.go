package models

import "time"

type Rate struct {
	ID           int       `gorm:"primary_key";"AUTO_INCREMENT"`
	SellCurrency string    `gorm:"column:sell_currency"`
	BuyCurrency  string    `gorm:"column:buy_currency"`
	Rate         float64   `gorm:"column:rate"`
	CreatedOn    time.Time `gorm:"column:created_on"`
}

func (rate *Rate) TableName() string {
	return "rates"
}
