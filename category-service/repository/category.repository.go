package repository

import (
	"category-service/model"
	"category-service/request"
	"log"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Migrate() error
	AddCategory(request.RequestAdd) (model.Category, error)
	ListCategory() ([]model.Category, error)
	DetailCategory(int) (model.Category, error)
	UpdateCategory(request.RequestUpdateCategory) (model.Category, error)
}

type categoryRepo struct {
	DB *gorm.DB
}

// UpdateCategory implements CategoryRepo.
func (c categoryRepo) UpdateCategory(req request.RequestUpdateCategory) (data model.Category, err error) {
	return data, c.DB.Model(&data).Where("id = ? ", req.Id).Updates(model.Category{
		Category: req.Category,
	}).Error
}

// DetailCategory implements CategoryRepo.
func (c categoryRepo) DetailCategory(req int) (data model.Category, err error) {
	return data, c.DB.First(&data, "id = ?", req).Error
}

// ListCategory implements CategoryRepo.
func (c categoryRepo) ListCategory() (data []model.Category, err error) {
	return data, c.DB.Find(&data).Error
}

// AddCategory implements CategoryRepo.
func (c categoryRepo) AddCategory(req request.RequestAdd) (data model.Category, err error) {
	data = model.Category{
		Category: req.Category,
	}

	return data, c.DB.Create(&data).Error
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return categoryRepo{
		DB: db,
	}
}

func (c categoryRepo) Migrate() error {
	log.Print("[CategoryRepository]...Migrate")
	return c.DB.AutoMigrate(&model.Category{})
}
