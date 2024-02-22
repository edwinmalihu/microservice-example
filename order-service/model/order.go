package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
}

func (Customer) TableName() string {
	return "customer"
}

type Order struct {
	gorm.Model
	CustomerID  uint      `json:"customer_id"`
	Customer    Customer  `gorm:"foreignKey:CustomerID"`
	TotalAmount float64   `json:"total_amount" gorm:"type:decimal(22,2)"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
}

func (Order) TableName() string {
	return "order"
}

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255);unique;not null"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Price       float64 `json:"price" gorm:"type:decimal(22,2)"`
	Stok        uint    `json:"stok"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `json:"qty"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

func (Product) TableName() string {
	return "product"
}
