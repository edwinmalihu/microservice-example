package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string `json:"category" gorm:"type:varchar(255);unique;not null"`
}

func (Category) TableName() string {
	return "category"
}
