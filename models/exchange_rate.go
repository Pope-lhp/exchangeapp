package models

import "time"

type ExchangeRate struct {
	ID           int       `gorm:"primary_key" json:"_id"`
	FromCurrency string    `json:"from_currency"binding:"required"`
	ToCurrency   string    `json:"to_currency"binding:"required"`
	Rate         float64   `json:"rate"binding:"required"`
	Date         time.Time `json:"date"`
} //交换货币的模型
