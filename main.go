package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api"
	"github.com/nathaliapavan/edu-api/infra/config"
)

func main() {
	gin.ForceConsoleColor()

	if err := config.StartConfig(); err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	if err := api.NewService().Start(); err != nil {
		log.Fatal(err)
	}
}
