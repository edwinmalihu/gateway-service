package route

import (
	"auth-services/controller"
	"auth-services/middleware"
	"auth-services/repository"
	"fmt"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	httpRouter.Use(middleware.CORSMiddleware())

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

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	customerRepo := repository.NewCustomerRepo()
	customerController := controller.NewCustomerController(customerRepo)
	apiRoutes := httpRouter.Group("/api")
	{
		apiRoutes.POST("/signin", customerController.Register)
		apiRoutes.POST("/register", customerController.Register)
	}

	//httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	httpRouter.Run(":8081")
}
