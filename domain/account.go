package domain

import "github.com/GURURAJ8/banking/errors"
import "github.com/GURURAJ8/banking/dto"
//This file contains the domain model and repository interface for Customer

type Account struct{
	Id 			string  `db:"account_id"`
	CustomerId 	int     `db:"customer_id"`	
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount 		float64 `db:"amount"`
	Status 		string  `db:"status"`
}

type AccountRepository interface{
	Save(a Account) (*Account, *errors.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		Account_id:  a.Id,
	}
}