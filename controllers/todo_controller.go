package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/config"
	"github.com/javadyousefi1/go-lang/models"
	"net/http"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Retrieve all todo items from the database
// @Tags Todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo "List of Todos"
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	result := config.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// CreateTodo godoc
// @Summary      Create a todo
// @Description  Create a new todo with title and status
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        todo  body  models.Todo  true  "Todo to create"
// @Success      201   {object}  models.Todo
// @Router       /todos [post]
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	// Validates based on binding tags
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	result := config.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
