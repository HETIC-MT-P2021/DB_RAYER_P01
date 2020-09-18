package model

import (
	"database/sql"
	"errors"

	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"
	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/entity"
)

// GetCustomerByNumber returns data of a customer.
func GetCustomerByNumber(customerNumber int) (e.Customer, error) {

	var customer e.Customer

	const query = `SELECT 
					c.customerNumber,
					c.customerName, 
					c.contactFirstName,
					c.contactLastName,
					c.phone,
					c.addressLine1,
					c.addressLine2,
					c.city,
					c.state,
					c.postalCode,
					c.country,
					c.salesRepEmployeeNumber,
					c.creditLimit,
					(SELECT GROUP_CONCAT(o.orderNumber)
						FROM orders o
						LEFT JOIN customers c
							ON c.customerNumber = o.customerNumber
						WHERE o.customerNumber = ?
					)
					FROM customers c
					WHERE c.customerNumber = ?`
	err := db.DB.QueryRow(query, customerNumber, customerNumber).Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.ContactLastName, &customer.ContactFirstName, &customer.Phone, &customer.AddressLine1, &customer.AddressLine2, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.SalesRepEmployeeNumber, &customer.CreditLimit, &customer.Order)

	if err == sql.ErrNoRows {
		return customer, errors.New("Customer is not found")
	}

	if err != nil {
		return customer, err
	}

	return customer, nil
}
