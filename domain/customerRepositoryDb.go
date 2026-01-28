package domain

import (
	"database/sql"
	"time"
"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/GURURAJ8/banking/errors"
)

type CustomerRepositoryDb struct {
	//db connection details would go here
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {		


FindAllSql := "SELECT customer_id, name, zipcode, city, status, date_of_birth FROM customers"
rows, err := c.client.Query(FindAllSql)
if err != nil {
log.Println("Error while querying customers table " + err.Error())
return nil, err
}
defer rows.Close()

var customers []Customer	
for rows.Next() {
	var customer Customer
	err := rows.Scan(&customer.Id, &customer.Name, &customer.Zip, &customer.City, &customer.Status, &customer.DateofBirth)
	if err != nil {
		return nil, err
	}
	customers = append(customers, customer)
}
return customers, nil
}

func (c CustomerRepositoryDb) ById(id int) (*Customer, *errors.AppError) {		
FindByIdSql := "SELECT customer_id, name, zipcode, city, status, date_of_birth FROM customers WHERE customer_id = ?"
row := c.client.QueryRow(FindByIdSql, id)
var customer Customer
err := row.Scan(&customer.Id, &customer.Name, &customer.Zip, &customer.City, &customer.Status, &customer.DateofBirth)
if err != nil {
	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundError("Customer not found")
	}else{
		log.Println("Error while querying customer by id " + err.Error())	
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
}
return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:SaRaPh11@tcp(localhost:3306)/banking")
if err != nil {
	panic(err)
}
// See "Important settings" section.
client.SetConnMaxLifetime(time.Minute * 3)
client.SetMaxOpenConns(10)
client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}