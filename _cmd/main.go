package main

import (
	"fmt"
	"log"
	"main/internal/adapter/config"
	"main/internal/adapter/storage/mongo"

	httpServer "main/_cmd/http"

	"github.com/fatih/color"
)

func main() {
	color.Green("----------------------> START SERVER <--------------------------")

	// load environment variables
	fmt.Println("load environment variables")
	config, err := config.New()
	if err != nil {
		log.Fatalf("Error loading environment variables", err)
	}

	// database
	fmt.Println("connecting to database")
	resource, err := mongo.New(config.DB)
	if err != nil {
		log.Fatalf("Connection database failure, Please check connection", err)
	}
	defer resource.Close()

	httpServer.HttpMain(config, resource)
}
