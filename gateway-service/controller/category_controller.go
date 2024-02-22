package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AddCategory(*gin.Context)
	UpdateCategory(*gin.Context)
	DetailCategory(*gin.Context)
	ListCategory(*gin.Context)
}

type categoryController struct {
	allRepo repository.CategoryRepo
}

// AddCategory	godoc
// @Summary Create Category
// @Description Create Category
// @param	signin body model.RequestAdd{} true "Signin"
// @Produce applicaton/json
// @Success 200 {object} model.ResponseCategory{}
// @Router /category/add [post]
// @Tags category-service
// AddCategory implements CategoryController.
func (r categoryController) AddCategory(ctx *gin.Context) {
	var req model.RequestAdd
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.AddCategory(req)
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

// DetailCategory	godoc
// @Summary Detail Category
// @Description Detail Category
// @Produce applicaton/json
// @Success 200 {object} model.ResponseCategory{}
// @Router /category/detail [get]
// @Param category_id query string true "Category ID"
// @Tags category-service
// DetailCategory implements CategoryController.
func (r categoryController) DetailCategory(ctx *gin.Context) {
	var req model.RequesListCategoryById
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.DetailCategory(req)
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

// ListCategory	godoc
// @Summary List Category
// @Description List Categor
// @Produce applicaton/json
// @Success 200 {object} []model.ResponseCategory{}
// @Router /category/list [get]
// @Tags category-service
// ListCategory implements CategoryController.
func (r categoryController) ListCategory(ctx *gin.Context) {

	data, resp, err := r.allRepo.ListCategory()
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

// UpdateCategory	godoc
// @Summary Update Category
// @Description Update Category
// @param	signin body model.RequestUpdateCategory{} true "Signin"
// @Produce applicaton/json
// @Success 200 {object} model.ResponseCategory{}
// @Router /category/update [post]
// @Tags category-service
// UpdateCategory implements CategoryController.
func (r categoryController) UpdateCategory(ctx *gin.Context) {
	var req model.RequestUpdateCategory
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.UpdateCategory(req)
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

func NewCategoryController(repo repository.CategoryRepo) CategoryController {
	return categoryController{
		allRepo: repo,
	}
}
