package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerId    uuid.UUID `json:"customerid" db:"customeruuid"`
	FirstName     string    `json:"firstname" db:"firstname"`
	MiddleInitial string    `json:"middleinitial" db:"middleinitial"`
	LastName      string    `json:"lastname" db:"lastname"`
}

type TempCustomer struct {
	CustomerId    uuid.UUID
	FirstName     sql.NullString
	MiddleInitial sql.NullString
	LastName      sql.NullString
}

type UpdateCustomer struct {
	CustomerId    uuid.UUID `json:"CustomerId" binding:"required" db:"customeruuid"`
	FirstName     *string   `json:"firstname,omitempty" db:"firstname"`
	MiddleInitial *string   `json:"middleinitial,omitempty" db:"middleinitial"`
	LastName      *string   `json:"lastname,omitempty" db:"lastname"`
}
