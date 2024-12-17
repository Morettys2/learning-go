package main

import (
	"github.com/Morettys2/learning-go/router" // Caminho correto do pacote router
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador
	r := gin.Default()

	// Configura as rotas
	router.InicializarRotas(r)

	// Executa o servidor na porta 8080
	r.Run(":8080") // Certifique-se de usar :8080
}
