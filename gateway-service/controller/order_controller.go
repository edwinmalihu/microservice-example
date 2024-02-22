package controller

import (
	"auth-services/model"
	"auth-services/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	AddOrder(*gin.Context)
}

type orderController struct {
	allRepo repository.OrderRepo
}

// CreateOrder	godoc
// @Summary Create Order
// @Description Add Product to Order
// @param	cary body model.AddOrder{} true "Create Order"
// @Produce applicaton/json
// @Success 200 {object} model.ResponsOrder{}
// @Router /api/order/add [post]
// @Tags order-service
// AddOrder implements OrderController.
func (r orderController) AddOrder(ctx *gin.Context) {
	var req model.AddOrder
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "informasi salah")
		return
	}

	data, resp, err := r.allRepo.AddOrder(req)
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

func NewOrderController(repo repository.OrderRepo) OrderController {
	return orderController{
		allRepo: repo,
	}
}
