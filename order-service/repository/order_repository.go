package repository

import (
	"log"
	"order-service/model"
	"order-service/request"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Migrate() error
	AddOrder(request.AddOrder) (model.Order, error)
	AddOrderItem(uint, uint, uint) (model.OrderItem, error)
}

type orderRepository struct {
	DB *gorm.DB
}

// AddOrder implements OrderRepository.
func (o orderRepository) AddOrder(req request.AddOrder) (data model.Order, err error) {
	data = model.Order{
		CustomerID:  req.CustomerID,
		TotalAmount: req.TotalAmount,
		OrderDate:   req.OrderDate,
		Status:      req.Status,
	}

	return data, o.DB.Create(&data).Error
}

// AddOrderItem implements OrderRepository.
func (o orderRepository) AddOrderItem(order_id uint, product_id uint, qty uint) (data model.OrderItem, err error) {
	data = model.OrderItem{
		OrderID:   order_id,
		ProductID: product_id,
		Quantity:  qty,
	}
	return data, o.DB.Create(&data).Error
}

// Migrate implements OrderRepository.
func (o orderRepository) Migrate() error {
	log.Print("[OrderRepository]...Migrate")
	return o.DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
}

func NewOrderRepo(db *gorm.DB) OrderRepository {
	return orderRepository{
		DB: db,
	}
}
