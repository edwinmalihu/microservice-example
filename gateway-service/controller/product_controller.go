package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"fmt"
	"net/http"

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
	allRepo repository.ProductRepo
}

// CreateProduct	godoc
// @Summary Create Product
// @Description Create Product
// @param	CreateProduct body model.RequestAddProduct{} true "Create Product"
// @Produce applicaton/json
// @Success 200 {object} model.ResponseSucessProduct{}
// @Router /product/add [post]
// @Tags product-service
// AddProduct implements ProductController.
func (r productController) AddProduct(ctx *gin.Context) {
	var req model.RequestAddProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.AddProduct(req)
	response := &model.Response{}
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

// DetailProduct	godoc
// @Summary Detail Product
// @Description Detail Product
// @Produce applicaton/json
// @Success 200 {object} model.ResponseDetailProduct{}
// @Router /product/detail [get]
// @Param product_id query string true "Product ID"
// @Tags product-service
// DetailProduct implements ProductController.
func (r productController) DetailProduct(ctx *gin.Context) {
	var req model.RequesByIdProduct
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.DetailProduct(req)
	response := &model.Response{}
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

// ListProduct	godoc
// @Summary List Product
// @Description List Product
// @Produce applicaton/json
// @Success 200 {object} model.ResponseDetailProduct{}
// @Router /product/list [get]
// @Tags product-service
// ListProduct implements ProductController.
func (r productController) ListProduct(ctx *gin.Context) {
	data, resp, err := r.allRepo.ListProduct()
	response := &model.Response{}
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

// ListProductByCategory	godoc
// @Summary List Product by Category
// @Description List Product by Category
// @Produce applicaton/json
// @Success 200 {object} model.ResponseDetailProduct{}
// @Router /product/listby-category [get]
// @Param category_id query string true "Category ID"
// @Tags product-service
// ListProductByCategory implements ProductController.
func (r productController) ListProductByCategory(ctx *gin.Context) {
	var req model.RequesByIdCategory
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.ListProductByCategory(req)
	response := &model.Response{}
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)

}

// UpdateProduct	godoc
// @Summary Update Product
// @Description Update Product
// @param	UpdateProduct body model.RequestUpdateProduct{} true "Update Product"
// @Produce applicaton/json
// @Success 200 {object} model.ResponseSucessProduct{}
// @Router /product/update [post]
// @Tags product-service
// UpdateProduct implements ProductController.
func (r productController) UpdateProduct(ctx *gin.Context) {
	var req model.RequestUpdateProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.UpdateProduct(req)
	response := &model.Response{}
	if err != nil {
		response.Success = false
		response.Status = http.StatusInternalServerError
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": err.Error()}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode() != 200 {
		response.Success = false
		response.Status = resp.StatusCode()
		response.ErrorCode = "9901"
		response.RespMessage = "Failed to get response"
		response.RespData = gin.H{"error": fmt.Sprint(resp)}
		ctx.JSON(resp.StatusCode(), response)
		return
	}

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

func NewProductController(repo repository.ProductRepo) ProductController {
	return productController{
		allRepo: repo,
	}
}
