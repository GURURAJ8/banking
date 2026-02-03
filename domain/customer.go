package domain

import "github.com/GURURAJ8/banking/errors"
import "github.com/GURURAJ8/banking/dto"
//This file contains the domain model and repository interface for Customer
type Customer struct{
	Id 	int `db:"customer_id"`
	Name string `db:"name"`
	Zip string `db:"zipcode"`
	City string `db:"city"`
	Status string `db:"status"`		
	DateofBirth string `db:"date_of_birth"`
}

//port interface
type CustomerRepository interface{
	// GetAllCustomers() []Customer
	FindAll() ([]Customer, error)
	ById(id int) (*Customer, *errors.AppError)	
}

func (c Customer) statusAsText() string {
	if c.Status == "1" {
		return "active"
	}
	return "inactive"
}

func (c *Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		CustomerId:  c.Id,
		Name:        c.Name,
		Zipcode:     c.Zip,
		City:        c.City,
		Status:      c.statusAsText(),
		DateOfBirth: c.DateofBirth,
}
}
