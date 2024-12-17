package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialziar() {
	// Cria um roteador Gin
	router := gin.Default()

	// Inicia o servidor na porta 8080
	router.Run(":8080") // Equivalente a r.Run(":8080")
}
