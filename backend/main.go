package main

import (
	"Expense_Tracker/db"
	"Expense_Tracker/models"
	"Expense_Tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Connect()
	db.DB.AutoMigrate(&models.Expense{})

	routes.Setup(r)
	r.Run(":8080")
}
