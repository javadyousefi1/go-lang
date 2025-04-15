package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javadyousefi1/go-lang/config"
	"github.com/javadyousefi1/go-lang/models"
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
