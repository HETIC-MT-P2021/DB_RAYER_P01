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

	const query = `SELECT *,
						(select concat('', string_agg(o.orderNumber, ','))
						FROM orders as o
								LEFT JOIN customers as c
										ON (c.customerNumber = o.customerNumber)
						WHERE customerNumber = $1) as orders
					FROM customers as c
					WHERE customerNumber = $1`
	err := db.DB.QueryRow(query, customerNumber).Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.ContactLastName, &customer.ContactFirstName, &customer.Phone, &customer.AddressLine1, &customer.AddressLine2, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.SalesRepEmployeeNumber, &customer.CreditLimit, &customer.Order)

	if err == sql.ErrNoRows {
		return customer, errors.New("Customer is not found")
	}

	if err != nil {
		return customer, err
	}

	return customer, nil
}