package main

import (
	"gingorm/config"
	"gingorm/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()

	if err == nil {
		logger.Errf("CONFIG INITIALIZATION ERROR: %v", err)
		return
	}

	router.Initialize()
}
