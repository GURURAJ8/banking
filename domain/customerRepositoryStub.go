package domain

type CustomerRepositoryStub struct{
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

//helper function to create a new stub repository with sample data
func NewCustomerRepositoryStub() CustomerRepositoryStub{
	customers := []Customer{
			{Id: 1, Name: "John Doe", Zip: "12345", City: "New York", Status: "Active", DateofBirth: "1990-01-01"},	{Id: 2, Name: "Jane Smith", Zip: "67890", City: "Los Angeles", Status: "Inactive", DateofBirth: "1985-05-15"},{Id: 3, Name: "Alice Johnson", Zip: "54321", City: "Chicago", Status: "Active", DateofBirth: "1992-09-30"},
	}
	return CustomerRepositoryStub{customers}
}