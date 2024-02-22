package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"auth-services/utils"
	"fmt"
	"net/http"

	_ "auth-services/docs"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type customerController struct {
	allRepo repository.CustomerRepository
}

// Signin	godoc
// @Summary Signin
// @Description Authentication and Unauthorized User
// @param	signin body model.Login{} true "Signin"
// @Produce applicaton/json
// @Success 200 {object} model.Login{}
// @Router /signin/ [post]
// @Tags customer-service
// Login implements CustomerController.
func (r customerController) Login(ctx *gin.Context) {
	var req model.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.Login(req)
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

	token := utils.GenerateToken(data.Username)
	data.Token = token

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

// Register	godoc
// @Summary Register
// @Description Create User
// @param	signin body model.AddCustomer{} true "Signin"
// @Produce applicaton/json
// @Success 200 {object} model.SuccessAddCustomer{}
// @Router /register/ [post]
// @Tags customer-service
// LoginCustomer implements CustomerController.
func (r customerController) Register(ctx *gin.Context) {
	var req model.AddCustomer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.Register(req)
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

	data.Password = ""

	response.Success = true
	response.Status = http.StatusOK
	response.ErrorCode = ""
	response.RespMessage = "Success to get response"
	response.RespData = data
	ctx.JSON(http.StatusOK, response)
}

func NewCustomerController(repo repository.CustomerRepository) CustomerController {
	return customerController{
		allRepo: repo,
	}
}
