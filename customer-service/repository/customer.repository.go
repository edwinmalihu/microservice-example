package repository

import (
	"customer-service/model"
	"customer-service/request"
	"log"

	"gorm.io/gorm"
)

type CustomerRepo interface {
	Migrate() error
	AddCustomer(request.AddCustomer) (model.Customer, error)
	UserLogin(request.Login) (model.Customer, error)
}

type customerRepo struct {
	DB *gorm.DB
}

// UserLogin implements CustomerRepo.
func (c customerRepo) UserLogin(req request.Login) (data model.Customer, err error) {
	return data, c.DB.First(&data, "username=?", req.Username).Error
}

// AddCustomer implements CustomerRepo.
func (c customerRepo) AddCustomer(req request.AddCustomer) (data model.Customer, err error) {
	data = model.Customer{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Email,
	}

	return data, c.DB.Create(&data).Error
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	return customerRepo{
		DB: db,
	}
}

func (c customerRepo) Migrate() error {
	log.Print("[CustomerRepository]...Migrate")
	return c.DB.AutoMigrate(&model.Customer{})
}
