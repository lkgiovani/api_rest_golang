package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	//inicializar o router utilizado as configurações default do gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//roda a API
	router.Run(":8000")
}