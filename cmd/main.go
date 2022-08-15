package main

import (
	"ascii-art-web/internal/controller"
	"ascii-art-web/internal/delivery"
	"log"
	"os"
)

func main() {
	config := os.Getenv("ASCII_WEB_CONFIGFILE")
	if config == "" {
		config = "config/config.json"
	}
	generator, err := controller.ReadConfigFile(config)
	if err != nil {
		log.Print(err)
		return
	}
	err = delivery.StartServer(generator)
	if err != nil {
		log.Print(err)
		return
	}
}
