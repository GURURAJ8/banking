package domain

//This file contains the domain model and repository interface for Customer
type Customer struct{
	Id 	int
	Name string
	Zip string
	City string
	Status string		
	DateofBirth string
}

type CustomerRepository interface{
	GetAllCustomers() []Customer
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct{
	customers []Customer
}

func (c CustomerRepositoryStub) GetAllCustomers() []Customer {
	return c.customers
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub{
	return CustomerRepositoryStub{
		customers: []Customer{
			{Id: 1, Name: "John Doe", Zip: "12345", City: "New York", Status: "Active", DateofBirth: "1990-01-01"},	{Id: 2, Name: "Jane Smith", Zip: "67890", City: "Los Angeles", Status: "Inactive", DateofBirth: "1985-05-15"},{Id: 3, Name: "Alice Johnson", Zip: "54321", City: "Chicago", Status: "Active", DateofBirth: "1992-09-30"},
		},
	}
}