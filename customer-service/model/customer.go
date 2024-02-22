package model

import "gorm.io/gorm"

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
