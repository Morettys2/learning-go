package main

import (
	"github.com/Morettys2/learning-go/config"
	"github.com/Morettys2/learning-go/router" // Caminho correto do pacote router
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
