package model

import (
	db "github.com/HETIC-MT-P2021/DB_RAYER_P01/database"
	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/entity"
)

// GetProductByOrderNumber returns data of a customer.
func GetProductByOrderNumber(orderNumber int) ([]e.Product, error) {

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
		var product e.Product
		err = rows.Scan(&product.ProductCode, &product.ProductName, &product.ProductLine, &product.ProductScale, &product.ProductVendor, &product.ProductDescription, &product.QuantityInStock, &product.BuyPrice, &product.MSRP)

		if err != nil {
			return productList, err
		}

		productList = append(productList, product)
	}

	err = rows.Err()

	if err != nil {
		return productList, err
	}

	return productList, nil
}
