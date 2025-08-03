package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/latestcomment/go-gin-web-server/internal/db"
	"github.com/latestcomment/go-gin-web-server/internal/models"
	"github.com/latestcomment/go-gin-web-server/internal/utils"
)

func GetAllCustomers() ([]models.Customer, error) {
	rows, err := db.DB.Query("select customeruuid, firstname, middleinitial, lastname from customers limit 5")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []models.Customer
	var tempCustomer models.TempCustomer
	for rows.Next() {
		scan_err := rows.Scan(&tempCustomer.CustomerId, &tempCustomer.FirstName, &tempCustomer.MiddleInitial, &tempCustomer.LastName)
		if scan_err != nil {
			return nil, scan_err
		}

		customers = append(customers, models.Customer{
			CustomerId:    tempCustomer.CustomerId,
			FirstName:     utils.NullToString(tempCustomer.FirstName),
			MiddleInitial: utils.NullToString(tempCustomer.MiddleInitial),
			LastName:      utils.NullToString(tempCustomer.LastName),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func GetCustomerByID(id uuid.UUID) (*models.Customer, error) {
	row := db.DB.QueryRow("SELECT customeruuid, firstname, middleinitial, lastname FROM customers WHERE customeruuid = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var tempCustomer models.TempCustomer

	if err := row.Scan(&tempCustomer.CustomerId, &tempCustomer.FirstName, &tempCustomer.MiddleInitial, &tempCustomer.LastName); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	customer := &models.Customer{
		CustomerId:    tempCustomer.CustomerId,
		FirstName:     utils.NullToString(tempCustomer.FirstName),
		MiddleInitial: utils.NullToString(tempCustomer.MiddleInitial),
		LastName:      utils.NullToString(tempCustomer.LastName),
	}
	return customer, nil
}

func UpdateCustomerByID(update models.UpdateCustomer) (err error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	setClauses := []string{}
	args := []interface{}{}

	i := 1
	if update.FirstName != nil {
		setClauses = append(setClauses, fmt.Sprintf(`firstname = $%d`, i))
		args = append(args, *update.FirstName)
		i++
	}
	if update.MiddleInitial != nil {
		setClauses = append(setClauses, fmt.Sprintf(`middleinitial = $%d`, i))
		args = append(args, *update.MiddleInitial)
		i++
	}
	if update.LastName != nil {
		setClauses = append(setClauses, fmt.Sprintf(`lastname = $%d`, i))
		args = append(args, *update.LastName)
		i++
	}

	args = append(args, update.CustomerId)
	query := fmt.Sprintf(
		`UPDATE customers
             SET %s
             WHERE customeruuid = $%d`,
		strings.Join(setClauses, ", "),
		i,
	)

	_, err = tx.Exec(query, args...)

	return err
}
