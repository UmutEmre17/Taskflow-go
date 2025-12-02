package main

import (
	"taskflow/models"
	"taskflow/routes"
	"github.com/gin-gonic/gin"
	"taskflow/config"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Task{})

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}