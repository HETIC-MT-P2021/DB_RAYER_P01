package model

import (
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"
	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/entity"
)

// GetProductByOrderNumber returns data of a customer.
func GetProductByOrderNumber(orderNumber int) ([]e.Product, error) {

	var product e.Product
	var productList []e.Product

	const query = `SELECT
					p.productCode,
					p.productName,
					p.productLine,
					p.productScale,
					p.productVendor,
					p.productDescription,
					p.quantityInStock,
					p.buyPrice,
					p.MSRP
					FROM orderdetails as od
					LEFT JOIN products as p
							ON (od.productCode = p.productCode)
					WHERE od.orderNumber = ?
				`

	rows, err := db.DB.Query(query, orderNumber)

	if err != nil {
		return productList, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&product.ProductCode, &product.ProductName, &product.ProductLine, &product.ProductScale, &product.ProductVendor, &product.ProductDescription, &product.QuantityInStock, &product.BuyPrice, &product.MSRP)

		if err != nil {
			return productList, err
		}

		productList = append(productList, product)

		product = e.Product{}
	}

	err = rows.Err()

	if err != nil {
		return productList, err
	}

	return productList, nil
}