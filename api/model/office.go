package model

import (
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"
	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/entity"
)

// GetAllOffice returns list of data of all office.
func GetAllOffice() ([]e.Office, error) {

	var officeList []e.Office

	rows, err := db.DB.Query(`SELECT 
							o.officeCode,
							o.city,
							o.phone,
							o.addressLine1,
							o.addressLine2,
							o.state,
							o.country,
							o.postalCode,
							o.territory,
							GROUP_CONCAT(e.employeeNumber)
							FROM offices o
							LEFT JOIN employees e
								ON e.officeCode = o.officeCode
							GROUP BY o.officeCode`)

	if err != nil {
		return officeList, err
	}
	defer rows.Close()

	for rows.Next() {
		var office e.Office
		err = rows.Scan(&office.OfficeCode, &office.City, &office.Phone, &office.AddressLine1, &office.AddressLine2, &office.State, &office.Country, &office.PostalCode, &office.Territory, &office.Employee)

		if err != nil {
			return officeList, err
		}

		officeList = append(officeList, office)
	}

	err = rows.Err()

	if err != nil {
		return officeList, err
	}

	return officeList, nil
}
