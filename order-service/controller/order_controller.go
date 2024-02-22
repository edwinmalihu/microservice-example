package controller

import (
	"log"
	"net/http"
	"order-service/repository"
	"order-service/request"
	"order-service/response"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	AddOrder(*gin.Context)
}

type orderController struct {
	orderRepo repository.OrderRepository
}

// AddOrder implements OrderController.
func (o orderController) AddOrder(ctx *gin.Context) {
	var req request.AddOrder
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("req: ", req)

	req.Status = "waiting payment"
	req.OrderDate = time.Now()
	data, err := o.orderRepo.AddOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dataItem, err := o.orderRepo.AddOrderItem(data.ID, req.ProductID, req.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(dataItem)

	res := response.ResponsOrder{
		OrderID: data.ID,
		Msg:     "Pesanan Success, Mohon Segera Selasaikan Pembayaran",
	}

	ctx.JSON(http.StatusOK, res)
}

func NewOrderController(repo repository.OrderRepository) OrderController {
	return orderController{
		orderRepo: repo,
	}
}
