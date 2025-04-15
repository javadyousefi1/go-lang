package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/controllers"
)

func RegisterTodoRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)
}

