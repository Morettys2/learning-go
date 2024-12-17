package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialziar() {
	// Cria um roteador Gin
	r := gin.Default()

	// Define uma rota GET
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Inicia o servidor na porta 8080
	r.Run() // Equivalente a r.Run(":8080")
}
