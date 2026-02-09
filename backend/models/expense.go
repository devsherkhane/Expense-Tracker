package models

type Expense struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}
