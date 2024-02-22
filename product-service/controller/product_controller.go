package controller

import (
	"log"
	"net/http"
	"product-service/repository"
	"product-service/request"
	"product-service/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DetailProduct(*gin.Context)
	ListProduct(*gin.Context)
	ListProductByCategory(*gin.Context)
}

type productController struct {
	prodRepo repository.ProductRepo
}

// ListProductByCategory implements ProductController.
func (p productController) ListProductByCategory(ctx *gin.Context) {
	var req request.RequesByIdCategory
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(req.Id)
	data, err := p.prodRepo.ListProductByCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// ListProduct implements ProductController.
func (p productController) ListProduct(ctx *gin.Context) {
	data, err := p.prodRepo.ListProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)

}

// DetailProduct implements ProductController.
func (p productController) DetailProduct(ctx *gin.Context) {
	var req request.RequesByIdProduct
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(req.Id)
	data, err := p.prodRepo.DetailProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// UpdateProduct implements ProductController.
func (p productController) UpdateProduct(ctx *gin.Context) {
	var req request.RequestUpdateProduct

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := p.prodRepo.UpdateProduct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Log Update Product : ", data)
	product_category, err := p.prodRepo.UpdateProductCateogry(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Log Update Product Category : ", product_category)

	result, err := p.prodRepo.DetailProduct(int(req.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := response.ResponseSucessProduct{
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		Stok:        result.Stok,
		Category:    result.Category,
		Msg:         "Data Berhasil di Ubah",
	}

	ctx.JSON(http.StatusOK, res)
}

// AddProduct implements ProductController.
func (p productController) AddProduct(ctx *gin.Context) {
	var req request.RequestAddProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := p.prodRepo.AddProduct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	relasi, err := p.prodRepo.AddProductCategory(req, data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Log Relasi : ", relasi)

	result, err := p.prodRepo.DetailProduct(int(data.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := response.ResponseSucessProduct{
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		Stok:        result.Stok,
		Category:    result.Category,
		Msg:         "Data Berhasil di Tambahkan",
	}

	ctx.JSON(http.StatusOK, res)

}

func NewProductRepo(repo repository.ProductRepo) ProductController {
	return productController{
		prodRepo: repo,
	}
}
