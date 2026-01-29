package domain

import (
	"database/sql"
	"time"
// "log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/GURURAJ8/banking/errors"
	"github.com/GURURAJ8/banking/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	//db connection details would go here
	client *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {		


FindAllSql := "SELECT customer_id, name, zipcode, city, status, date_of_birth FROM customers"
var customers []Customer
err := c.client.Select(&customers, FindAllSql)
if err != nil {
logger.Error("Error while querying customers table " + err.Error())
return nil, err
}
// defer rows.Close()
	

// err = sqlx.StructScan(rows, &customers)
// if err != nil {
// 	logger.Error("Error while scanning customers " + err.Error())	
// 	return nil, err
// }
// for rows.Next() {
// 	var customer Customer
// 	err := rows.Scan(&customer.Id, &customer.Name, &customer.Zip, &customer.City, &customer.Status, &customer.DateofBirth)
// 	if err != nil {
// 		return nil, err
// 	}
// 	customers = append(customers, customer)
// }
return customers, nil
}

func (c CustomerRepositoryDb) ById(id int) (*Customer, *errors.AppError) {		
FindByIdSql := "SELECT customer_id, name, zipcode, city, status, date_of_birth FROM customers WHERE customer_id = ?"
// row := c.client.QueryRow(FindByIdSql, id)
var customer Customer
err := c.client.Get(&customer, FindByIdSql, id)
if err != nil {
	if err == sql.ErrNoRows {}
		return nil, errors.NewNotFoundError("Customer not found")
	}else{
		logger.Error("Error while querying customer by id " + err.Error())	
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

// err := row.Scan(&customer.Id, &customer.Name, &customer.Zip, &customer.City, &customer.Status, &customer.DateofBirth)
// if err != nil {
// 	if err == sql.ErrNoRows {
// 		return nil, errors.NewNotFoundError("Customer not found")
// 	}else{
// 		logger.Error("Error while querying customer by id " + err.Error())	
// 		return nil, errors.NewUnexpectedError("Unexpected database error")
// 	}
// }
return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:SaRaPh11@tcp(localhost:3306)/banking")
if err != nil {
	panic(err)
}
// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}