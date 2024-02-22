package route

import (
	"customer-service/controller"
	"customer-service/middleware"
	"customer-service/repository"
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(middleware.CORSMiddleware())

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	enforcer.EnableAutoSave(true)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	customerRepository := repository.NewCustomerRepo(db)

	if err := customerRepository.Migrate(); err != nil {
		log.Fatal("Customer migrate err", err)
	}

	customerController := controller.CustomerNewController(customerRepository)

	apiRoutes := httpRouter.Group("/api")

	{
		apiRoutes.POST("/add", customerController.AddCustomer(enforcer))
		apiRoutes.POST("/login", customerController.LoginUser)
	}

	httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	//httpRouter.Run(":8082")
}
