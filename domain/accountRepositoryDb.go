package domain

import "github.com/GURURAJ8/banking/errors"
// import "github.com/GURURAJ8/banking/dto"
import "github.com/jmoiron/sqlx"
import "log"

//This file contains the domain model and repository interface for Customer
type AccountRepositoryDb struct{
	//db sql.DB
	client *sqlx.DB
}

func (r AccountRepositoryDb) Save(a Account) (*Account, *errors.AppError){
	insertSql := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := r.client.Exec(insertSql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		log.Panic("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error occurred while creating new account")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panic("Error while getting last insert id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error occurred while creating new account")
	}
	a.Id = string(id)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}