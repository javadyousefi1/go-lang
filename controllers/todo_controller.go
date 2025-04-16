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

// DeleteTodo godoc
// @Summary Delete a todo by ID
// @Description Delete a todo item from the database by its ID
// @Tags Todos
// @Accept  json
// @Produce  json
// @Param   id   path    string  true  "Todo ID"  // ID of the todo to delete
// @Success 204  {string} string  "No Content"
// @Failure 400  {object} models.ErrorResponse  "Bad Request"
// @Failure 404  {object} models.ErrorResponse  "Not Found"
// @Failure 500  {object} models.ErrorResponse  "Internal Server Error"
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// Try to find the todo with the given ID
	var todo models.Todo
	result := config.DB.First(&todo, "id = ?", id)

	if result.Error != nil {
		// If not found, return a 404 error
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Delete the todo from the database
	result = config.DB.Delete(&todo)

	if result.Error != nil {
		// If an error occurs during deletion, return a 500 error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Successfully deleted, return a 204 No Content
	c.Status(http.StatusOK)
}
