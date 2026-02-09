package routes

import (
	"Expense_Tracker/db"
	"Expense_Tracker/models"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/expenses", getExpenses)
	r.POST("/expenses", addExpense)
	r.DELETE("/expenses/:id", deleteExpense)
}

func getExpenses(c *gin.Context) {
	var expenses []models.Expense
	db.DB.Find(&expenses)

	c.JSON(200, expenses)
}

func addExpense(c *gin.Context) {
	var expense models.Expense

	if err := c.BindJSON(&expense); err != nil {
		return
	}

	db.DB.Create(&expense)
	c.JSON(200, expense)
}

func deleteExpense(c *gin.Context) {
	id := c.Param("id")

	db.DB.Delete(&models.Expense{}, id)
	c.JSON(200, gin.H{"message": "Deleted"})
}
