package main

import (
	"exchangeapp/config" //  模块名/路径
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port) //

}
