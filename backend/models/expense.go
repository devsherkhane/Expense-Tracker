package models

import "time"

type Expense struct {
    // Change 'id' to 'ID' so Go exports it
	ID        uint      `json:"id" gorm:"primaryKey"` 
	Title     string    `json:"title" binding:"required,min=2"`
	Amount    float64   `json:"amount" binding:"required,gt=0"`
	Category  string    `json:"category" binding:"required"`
	Date      string    `json:"date" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}