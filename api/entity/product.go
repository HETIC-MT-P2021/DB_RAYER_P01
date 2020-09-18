package entity

// Product defines the structure of the product entity.
type Product struct {
	ProductCode        string  `json:"productCode"`
	ProductName        string  `json:"productName"`
	ProductLine        string  `json:"productLine"`
	ProductScale       string  `json:"productScale"`
	ProductVendor      string  `json:"productVendor"`
	ProductDescription string  `json:"productDescription"`
	QuantityInStock    int     `json:"quantityInStock"`
	BuyPrice           float64 `json:"buyPrice"`
	MSRP               float32 `json:"MSRP"`
}
