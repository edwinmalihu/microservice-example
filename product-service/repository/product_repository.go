package repository

import (
	"log"

	"product-service/model"
	"product-service/request"
	"product-service/response"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Migrate() error
	AddProduct(request.RequestAddProduct) (model.Product, error)
	AddProductCategory(request.RequestAddProduct, uint) (model.Product_Category, error)
	UpdateProduct(request.RequestUpdateProduct) (model.Product, error)
	UpdateProductCateogry(request.RequestUpdateProduct) (model.Product_Category, error)
	DetailProduct(int) (response.ResponseDetailProduct, error)
	ListProduct() ([]response.ResponseDetailProduct, error)
	ListProductByCategory(int) ([]response.ResponseDetailProduct, error)
}

type productRepo struct {
	DB *gorm.DB
}

// ListProductByCategory implements ProductRepo.
func (p productRepo) ListProductByCategory(req int) (data []response.ResponseDetailProduct, err error) {
	return data, p.DB.Raw("select p.name, p.description, p.price, p.stok, c.category from product as p join product_category as pc on p.id = pc.product_id join category as c on pc.category_id = c.id where c.id = ?", req).Scan(&data).Error
}

// ListProduct implements ProductRepo.
func (p productRepo) ListProduct() (data []response.ResponseDetailProduct, err error) {
	return data, p.DB.Raw("select p.name, p.description, p.price, p.stok, c.category from product as p join product_category as pc on p.id = pc.product_id join category as c on pc.category_id = c.id").Scan(&data).Error
}

// DetailProduct implements ProductRepo.
func (p productRepo) DetailProduct(req int) (data response.ResponseDetailProduct, err error) {
	return data, p.DB.Raw("select p.name, p.description, p.price, p.stok, c.category from product as p join product_category as pc on p.id = pc.product_id join category as c on pc.category_id = c.id where p.id = ?", req).Scan(&data).Error
}

// UpdateProductCateogry implements ProductRepo.
func (p productRepo) UpdateProductCateogry(req request.RequestUpdateProduct) (data model.Product_Category, err error) {
	return data, p.DB.Model(&data).Where("product_id = ?", req.Id).Updates(model.Product_Category{
		CategoryID: req.CategoryID,
	}).Error
}

// UpdateProduct implements ProductRepo.
func (p productRepo) UpdateProduct(req request.RequestUpdateProduct) (data model.Product, err error) {
	return data, p.DB.Model(&data).Where("id = ?", req.Id).Updates(model.Product{
		Name:        req.Name,
		Description: req.Description,
		Stok:        req.Stok,
		Price:       req.Price,
	}).Error
}

// AddProductCategory implements ProductRepo.
func (p productRepo) AddProductCategory(req request.RequestAddProduct, id uint) (data model.Product_Category, err error) {
	data = model.Product_Category{
		ProductID:  id,
		CategoryID: req.CategoryID,
	}

	return data, p.DB.Create(&data).Error
}

// AddProduct implements ProductRepo.
func (p productRepo) AddProduct(req request.RequestAddProduct) (data model.Product, err error) {
	data = model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stok:        req.Stok,
	}

	return data, p.DB.Create(&data).Error
}

func (p productRepo) Migrate() error {
	log.Print("[ProductRepository]...Migrate")
	return p.DB.AutoMigrate(&model.Product{})
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return productRepo{
		DB: db,
	}
}
