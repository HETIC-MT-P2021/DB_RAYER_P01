package entity

import (
	"gopkg.in/guregu/null.v4"
)

// Customer defines the structure of the customer entity.
type Customer struct {
	CustomerNumber         int         `json:"customerNumber"`
	CustomerName           string      `json:"customerName"`
	ContactLastName        string      `json:"contactLastName"`
	ContactFirstName       string      `json:"contactFirstName"`
	Phone                  string      `json:"phone"`
	AddressLine1           string      `json:"adressLine1"`
	AddressLine2           null.String `json:"adressLine2"`
	City                   string      `json:"city"`
	State                  null.String `json:"state"`
	PostalCode             null.String `json:"postalCode"`
	Country                string      `json:"country"`
	SalesRepEmployeeNumber null.Int    `json:"salesRepEmployeeNumber"`
	CreditLimit            null.Float  `json:"creditLimit"`
	Order                  null.String `json:"orders"`
	Product                null.String `json:"products"`
	NumberProduct          null.Int    `json:"numberProducts"`
	TotalPrice             null.Float  `json:"totalPrice"`
}
