package entity

import (
	d "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"
)

// Customer defines the structure of the customer entity.
type Customer struct {
	CustomerNumber         int           `json:"customerNumber"`
	CustomerName           string        `json:"customerName"`
	ContactLastName        string        `json:"contactLastName"`
	ContactFirstName       string        `json:"contactFirstName"`
	Phone                  string        `json:"phone"`
	AddressLine1           string        `json:"adressLine1"`
	AddressLine2           d.NullString  `json:"adressLine2"`
	City                   string        `json:"city"`
	State                  d.NullString  `json:"state"`
	PostalCode             d.NullString  `json:"postalCode"`
	Country                string        `json:"country"`
	SalesRepEmployeeNumber d.NullInt64   `json:"salesRepEmployeeNumber"`
	CreditLimit            d.NullFloat64 `json:"creditLimit"`
	Order                  d.NullString  `json:"order"`
}
