package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/config"
	"github.com/javadyousefi1/go-lang/models"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	result := config.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

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
