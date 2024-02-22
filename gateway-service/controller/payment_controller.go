package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	AddPayment(*gin.Context)
	DeletePayment(*gin.Context)
}

type paymentController struct {
	allRepo repository.PaymentRepo
}

// CreatePayment	godoc
// @Summary Create Paymenr
// @Description Add order to payment
// @param	cary body model.RequestAddPayment{} true "Payment Order"
// @Produce applicaton/json
// @Success 200 {object} model.ResponseSuccessPayment{}
// @Router /api/payment/add [post]
// @Tags payment-service
// AddPayment implements PaymentController.
func (r paymentController) AddPayment(ctx *gin.Context) {
	var req model.RequestAddPayment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.AddPayment(req)
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

// DeletePayment	godoc
// @Summary Delete Payment
// @Description Delete Payment Order
// @Produce applicaton/json
// @Success 200 {object} model.Response{}
// @Router /api/payment/delete [delete]
// @Param payment_id query string true "Payment ID"
// @Tags payment-service
// DeletePayment implements PaymentController.
func (r paymentController) DeletePayment(ctx *gin.Context) {
	var req model.RequestByIdPayment
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.DeletePayment(req)
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

func NewPaymentController(repo repository.PaymentRepo) PaymentController {
	return paymentController{
		allRepo: repo,
	}
}
