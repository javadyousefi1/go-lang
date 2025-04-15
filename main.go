package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/config"
	"github.com/javadyousefi1/go-lang/models"
	"github.com/javadyousefi1/go-lang/routes"
)

func main() {
	config.InitConfig()
	config.DB.AutoMigrate(&models.Todo{})

	router := gin.Default()
	routes.RegisterTodoRoutes(router)

	router.Run(":8080")
}
