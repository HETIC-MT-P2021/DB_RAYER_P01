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
					) as orders,
					(SELECT GROUP_CONCAT(od.productCode)
						FROM orders as o
							LEFT JOIN orderdetails AS od
								ON c.customerNumber = o.customerNumber
								AND od.orderNumber = o.orderNumber
						WHERE c.customerNumber = ?
					) as products,
					(SELECT COUNT(od.productCode)
					FROM orders as o
						LEFT JOIN orderdetails AS od
							ON c.customerNumber = o.customerNumber
							AND od.orderNumber = o.orderNumber
						WHERE c.customerNumber = ?
					) as number_products,
					(SELECT SUM(od.priceEach)
					FROM orders as o
						LEFT JOIN orderdetails AS od
							ON c.customerNumber = o.customerNumber
							AND od.orderNumber = o.orderNumber
						WHERE c.customerNumber = ?
					) as total_price
					FROM customers c
					WHERE c.customerNumber = ?`
	err := db.DB.QueryRow(query, customerNumber, customerNumber, customerNumber, customerNumber, customerNumber).Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.ContactLastName, &customer.ContactFirstName, &customer.Phone, &customer.AddressLine1, &customer.AddressLine2, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.SalesRepEmployeeNumber, &customer.CreditLimit, &customer.Order, &customer.Product, &customer.NumberProduct, &customer.TotalPrice)

	if err == sql.ErrNoRows {
		return customer, errors.New("Customer is not found")
	}

	if err != nil {
		return customer, err
	}

	return customer, nil
}
