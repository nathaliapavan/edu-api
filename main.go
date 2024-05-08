package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathaliapavan/edu-api/api"
)

func main() {
	gin.ForceConsoleColor()
	service := api.NewService()
	service.Start()
}
