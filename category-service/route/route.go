package route

import (
	"category-service/controller"
	"category-service/middleware"
	"category-service/repository"
	"fmt"
	"log"
	"os"

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

	categoryRepository := repository.NewCategoryRepo(db)

	if err := categoryRepository.Migrate(); err != nil {
		log.Fatal("Category migrate err", err)
	}

	categoryController := controller.CategoryNewController(categoryRepository)

	apiRoutes := httpRouter.Group("/api")

	{
		apiRoutes.POST("/add", categoryController.AddCategory)
		apiRoutes.GET("/list", categoryController.ListCategory)
		apiRoutes.GET("/detail", categoryController.DetailCategory)
		apiRoutes.POST("/update", categoryController.UpdateCategory)
	}

	httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	// httpRouter.Run(":8085")
}
