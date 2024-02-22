package model

import "time"

type AddOrder struct {
	CustomerID  uint      `json:"customer_id"`
	TotalAmount float64   `json:"total_amount"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	ProductID   uint      `json:"product_id"`
	Quantity    uint      `json:"qty"`
}

type ResponsOrder struct {
	OrderID uint   `json:"order_id"`
	Msg     string `json:"msg"`
}
