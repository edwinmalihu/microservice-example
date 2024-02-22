package model

import "time"

type RequestAddPayment struct {
	OrderID     uint      `json:"order_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

type RequestByIdPayment struct {
	PaymentID string `form:"json:payment_id"`
}

type ResponseSuccessPayment struct {
	PaymentID   uint      `json:"payment_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

type ResponseDelete struct {
	Status string `json:"status"`
}
