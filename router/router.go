package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Initialize the Router used the GIN default settings
	
	router := gin.Default()
	
	// initialize routes
	initializeRoutes(router)

	//run router API
	router.Run(":8000")
}