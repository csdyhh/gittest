package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.Status = 1
	employee.CreateTime = time.Now()
	employee.UpdateTime = time.Now()

	if err := models.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create an employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Employee created successfully",
	})
}
