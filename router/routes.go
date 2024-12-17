package router

import (
	"github.com/Morettys2/learning-go/handler"
	"github.com/gin-gonic/gin"
)

// InicializarRotas define as rotas da API
func InicializarRotas(router *gin.Engine) {
	// Cria um grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Define a rota GET dentro do grupo
		v1.GET("/opening", handler.Show)
		// Define a rota POST
		v1.POST("/opening", handler.Create)
		// Define a rota DELETE
		v1.DELETE("/opening", handler.Delete)
		// Define a rota PUT
		v1.PUT("/opening", handler.Update)
		// Define a rota GET
		v1.GET("/opening", handler.List)
	}
}
