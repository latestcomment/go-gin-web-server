package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/latestcomment/go-gin-web-server/internal/models"
	"github.com/latestcomment/go-gin-web-server/internal/services"
)

func GetAllCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, customers)
}

func GetCustomerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := services.GetCustomerByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if customer == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, customer)
}

func UpdateCustomerByID(c *gin.Context) {
	var update models.UpdateCustomer

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	err := services.UpdateCustomerByID(update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}
