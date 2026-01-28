package domain

import "github.com/GURURAJ8/banking/errors"
//This file contains the domain model and repository interface for Customer
type Customer struct{
	Id 	int
	Name string
	Zip string
	City string
	Status string		
	DateofBirth string
}

//port interface
type CustomerRepository interface{
	// GetAllCustomers() []Customer
	FindAll() ([]Customer, error)
	ById(id int) (*Customer, *errors.AppError)	
}

