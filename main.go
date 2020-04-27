package main

import (
	"google-rtb/config"
	"google-rtb/pkg/logger"
	r "google-rtb/router"
)

func main() {
	config.LoadConfig()
	logger.Init()

	router := r.GetRouter()

	err := router.Run(r.GetPort())

	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to start service:", params)
	}
}
