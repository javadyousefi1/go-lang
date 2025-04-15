package models

// Todo represents a task in the todo list
// @Description Represents a task with title and completion status

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}
