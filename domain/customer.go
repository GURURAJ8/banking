package domain

import "github.com/GURURAJ8/banking/errors"
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

