package entity

// Customer defines the structure of the customer entity.
type Customer struct {
	CustomerNumber         int    `json:"customer_number"`
	CustomerName           string `json:"customer_name"`
	ContactLastName        string `json:"contact_last_name"`
	ContactFirstName       string `json:"contact_first_name"`
	Phone                  string `json:"phone"`
	AddressLine1           string `json:"adress_line_1"`
	AddressLine2           string `json:"adress_line_2"`
	City                   string `json:"city"`
	State                  string `json:"state"`
	PostalCode             string `json:"postal_code"`
	Country                string `json:"country"`
	SalesRepEmployeeNumber string `json:"sales_rep_employee_number"`
	CreditLimit            string `json:"credit_limit"`
	Order                  string `json:"order"`
}
