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
					GROUP_CONCAT(DISTINCT orders.orderNumber) as orders,
					GROUP_CONCAT(orders.productCode) as products,
					COUNT(orders.productCode) as countProducts,
					SUM(orders.priceEach) as totalPrice
			 	FROM customers AS c
			 	LEFT JOIN
				  (SELECT o.customerNumber, od.productCode, od.orderNumber, od.priceEach
				   FROM orders as o
							LEFT JOIN orderdetails AS od
									  ON od.orderNumber = o.orderNumber
				   GROUP BY o.customerNumber, od.productCode, od.orderNumber
				  ) as orders
				  ON c.customerNumber = orders.customerNumber
			 	WHERE c.customerNumber = ?
			 	GROUP BY c.customerNumber`
	err := db.DB.QueryRow(query, customerNumber).Scan(&customer.CustomerNumber, &customer.CustomerName, &customer.ContactLastName, &customer.ContactFirstName, &customer.Phone, &customer.AddressLine1, &customer.AddressLine2, &customer.City, &customer.State, &customer.PostalCode, &customer.Country, &customer.SalesRepEmployeeNumber, &customer.CreditLimit, &customer.Order, &customer.Product, &customer.NumberProduct, &customer.TotalPrice)

	if err == sql.ErrNoRows {
		return customer, errors.New("Customer is not found")
	}

	if err != nil {
		return customer, err
	}

	return customer, nil
}
