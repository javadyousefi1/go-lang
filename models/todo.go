package models

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}
