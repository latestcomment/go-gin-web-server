package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/latestcomment/go-gin-web-server/internal/models"
	"github.com/latestcomment/go-gin-web-server/internal/repositories"
)

func GetAllCustomers() ([]models.Customer, error) {
	customers, err := repositories.GetAllCustomers()
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	return customers, nil
}

func GetCustomerByID(id uuid.UUID) (*models.Customer, error) {
	customer, err := repositories.GetCustomerByID(id)
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	return customer, nil
}

func UpdateCustomerByID(update models.UpdateCustomer) error {
	err := repositories.UpdateCustomerByID(update)
	return err
}
