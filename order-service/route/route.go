package route

import (
	"fmt"
	"log"
	"order-service/controller"
	"order-service/middleware"
	"order-service/repository"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB) {
	httpRoute := gin.Default()
	httpRoute.Use(middleware.CORSMiddleware())

	httpRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Status": "IP"})
	})

	orderRepo := repository.NewOrderRepo(db)
	if err := orderRepo.Migrate(); err != nil {
		log.Fatal("Order & Order Item migrate err", err)
	}

	orderController := controller.NewOrderController(orderRepo)

	apiRoute := httpRoute.Group("/api")
	{
		apiRoute.POST("/add", orderController.AddOrder)
	}

	// httpRoute.Run(":8083")
	httpRoute.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}
