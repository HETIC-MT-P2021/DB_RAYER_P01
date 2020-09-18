package entity

import (
	"gopkg.in/guregu/null.v4"
)

// Employee defines the structure of the employee entity.
type Employee struct {
	EmployeeNumber     int64    `json:"employeeNumber"`
	EmployeeLastName   string   `json:"lastName"`
	EmployeeFirstName  string   `json:"firstName"`
	EmployeeExtension  string   `json:"extension"`
	EmployeeEmail      string   `json:"email"`
	EmployeeOfficeCode string   `json:"officeCode"`
	EmployeeReportTo   null.Int `json:"reportsTo"`
	EmployeeJobTitle   string   `json:"jobTitle"`
}
