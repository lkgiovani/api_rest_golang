package main

import (
	"github.com/lkgiovani/api_rest_golang/config"
	"github.com/lkgiovani/api_rest_golang/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize config

	err := config.Init()

	if err != nil {
		logger.ErrF("Config initialization error: %v", err)
		return
	}
	
	// initialize router
	router.InitializeRouter()
}