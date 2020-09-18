package entity

import (
	h "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/helper"
)

// Customer defines the structure of the customer entity.
type Customer struct {
	CustomerNumber         int           `json:"customerNumber"`
	CustomerName           string        `json:"customerName"`
	ContactLastName        string        `json:"contactLastName"`
	ContactFirstName       string        `json:"contactFirstName"`
	Phone                  string        `json:"phone"`
	AddressLine1           string        `json:"adressLine1"`
	AddressLine2           h.NullString  `json:"adressLine2"`
	City                   string        `json:"city"`
	State                  h.NullString  `json:"state"`
	PostalCode             h.NullString  `json:"postalCode"`
	Country                string        `json:"country"`
	SalesRepEmployeeNumber h.NullInt64   `json:"salesRepEmployeeNumber"`
	CreditLimit            h.NullFloat64 `json:"creditLimit"`
	Order                  h.NullString  `json:"orders"`
	Product                h.NullString  `json:"products"`
	NumberProduct          h.NullString  `json:"numberProducts"`
	TotalPrice             h.NullString  `json:"totalPrice"`
}
