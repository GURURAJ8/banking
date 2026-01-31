package domain

import (
	"database/sql"
	"time"
// "log"
	"os"
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
return customers, nil
}

func (c CustomerRepositoryDb) ById(id int) (*Customer, *errors.AppError) {		
FindByIdSql := "SELECT customer_id, name, zipcode, city, status, date_of_birth FROM customers WHERE customer_id = ?"
var customer Customer
err := c.client.Get(&customer, FindByIdSql, id)
if err != nil {
	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundError("Customer not found")
	}else{
		logger.Error("Error while querying customer by id " + err.Error())	
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

}
return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser :=os.Getenv("DB_USER")
	dbPass :=os.Getenv("DB_PASS")
	dbAddr :=os.Getenv("DB_ADDR")
	dbPort :=os.Getenv("DB_PORT")
	dbName :=os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbAddr+":"+dbPort+")/"+dbName)
if err != nil {
	panic(err)
}
// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}