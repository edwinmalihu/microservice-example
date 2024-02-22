package controller

import (
	"customer-service/repository"
	"customer-service/request"
	"customer-service/response"
	"customer-service/utils"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	AddCustomer(enforcer *casbin.Enforcer) gin.HandlerFunc
	LoginUser(*gin.Context)
}

type customerController struct {
	custRepo repository.CustomerRepo
}

func CustomerNewController(repo repository.CustomerRepo) customerController {
	return customerController{
		custRepo: repo,
	}
}

func (h customerController) AddCustomer(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.AddCustomer
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req.Password = fmt.Sprintf("%s%s", req.Username, req.Password)
		utils.HashPassword(&req.Password)

		user, err1 := h.custRepo.AddCustomer(req)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, err1.Error())
			return
		}

		res := response.SuccessAddCustomer{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (h customerController) LoginUser(ctx *gin.Context) {
	var req request.Login

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := h.custRepo.UserLogin(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Informasi yang anda masukan salah")
		return

	}

	password := fmt.Sprintf("%s%s", req.Username, req.Password)
	utils.HashPassword(&password)
	isTrue := utils.ComparePassword(dbUser.Password, password)
	if !isTrue {
		ctx.JSON(http.StatusInternalServerError, "Informasi yang anda masukan salah")
		return
	}

	res := response.LoginResponse{
		Username: dbUser.Username,
		Msg:      "Login Success",
	}

	ctx.JSON(http.StatusOK, res)
}
