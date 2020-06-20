package main

import (
	"gor/pkg/config"
	"gor/pkg/server"
)

func main() {
	configuration, configErr := config.Init()
	if configErr == nil {
		return
	}

	serverConfig := configuration.Server
	server.Start(serverConfig.Port)
}
