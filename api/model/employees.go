package model

import (
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/db"
	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/entity"
)

// GetAllEmployee returns list of data of all employees.
func GetAllEmployee() ([]e.Employee, error) {

	var employeeList []e.Employee

	rows, err := db.DB.Query(`SELECT 
							e.employeeNumber,
							e.lastName,
							e.firstName,
							e.extension,
							e.email,
							e.officeCode,
							e.reportsTo,
							e.jobTitle
							FROM employees as e
							ORDER BY e.employeeNumber`)

	if err != nil {
		return employeeList, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee e.Employee
		err = rows.Scan(&employee.EmployeeNumber, &employee.EmployeeLastName, &employee.EmployeeFirstName, &employee.EmployeeExtension, &employee.EmployeeEmail, &employee.EmployeeOfficeCode, &employee.EmployeeReportTo, &employee.EmployeeJobTitle)

		if err != nil {
			return employeeList, err
		}

		employeeList = append(employeeList, employee)
	}

	err = rows.Err()

	if err != nil {
		return employeeList, err
	}

	return employeeList, nil
}
