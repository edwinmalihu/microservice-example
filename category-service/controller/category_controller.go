package controller

import (
	"category-service/repository"
	"category-service/request"
	"category-service/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AddCategory(*gin.Context)
	ListCategory(*gin.Context)
	DetailCategory(*gin.Context)
	UpdateCategory(*gin.Context)
}

type categoryController struct {
	custRepo repository.CategoryRepo
}

// DetailCategory implements CategoryController.
func (c categoryController) DetailCategory(ctx *gin.Context) {
	var req request.RequesListCategoryById
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(req.Id)

	data, err := c.custRepo.DetailCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := response.ResponseCategory{
		Id:       data.ID,
		Category: data.Category,
	}

	ctx.JSON(http.StatusOK, res)
}

// ListCategory implements CategoryController.
func (c categoryController) ListCategory(ctx *gin.Context) {
	var result []response.ResponseCategory
	data, err := c.custRepo.ListCategory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for _, list := range data {
		list_data := response.ResponseCategory{
			Id:       list.ID,
			Category: list.Category,
		}

		result = append(result, list_data)
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdateCategory implements CategoryController.
func (c categoryController) UpdateCategory(ctx *gin.Context) {
	var req request.RequestUpdateCategory
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.custRepo.UpdateCategory(req)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	log.Println(data)

	result, err := c.custRepo.DetailCategory(req.Id)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res := response.ResponseSuccess{
		Category: result.Category,
		Msg:      "Data Berhasil di Update",
	}

	ctx.JSON(http.StatusOK, res)
}

// AddCategory implements CategoryController.
func (c categoryController) AddCategory(ctx *gin.Context) {
	var req request.RequestAdd
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.custRepo.AddCategory(req)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res := response.ResponseSuccess{
		Category: data.Category,
		Msg:      "Data Behasil diTambahkan",
	}

	ctx.JSON(http.StatusOK, res)
}

func CategoryNewController(repo repository.CategoryRepo) CategoryController {
	return categoryController{
		custRepo: repo,
	}
}
