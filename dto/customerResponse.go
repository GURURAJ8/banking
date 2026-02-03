package dto 

type CustomerResponse struct {
	CustomerId   int    `json:"customer_id"`
	Name         string `json:"name"`
	Zipcode      string `json:"zipcode"`
	City         string `json:"city"`
	Status       string `json:"status"`
	DateOfBirth  string `json:"date_of_birth"`
}