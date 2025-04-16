// @title Todo API
// @version 1.0
// @description Simple API to manage todos
// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/config"
	_ "github.com/javadyousefi1/go-lang/docs"
	"github.com/javadyousefi1/go-lang/models"
	"github.com/javadyousefi1/go-lang/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.InitConfig()
	config.DB.AutoMigrate(&models.Todo{})

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterTodoRoutes(router)

	router.Run(":8080")
}
