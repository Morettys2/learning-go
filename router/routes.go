package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InicializarRotas define as rotas da API
func InicializarRotas(router *gin.Engine) {
	// Cria um grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Define a rota GET dentro do grupo
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "GET Ok"})
		})

		// Define a rota POST
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "POST Ok"})
		})

		// Define a rota DELETE
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "DELETE Ok"})
		})

		// Define a rota PUT
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PUT Ok"})
		})
	}
}
