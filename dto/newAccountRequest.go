package dto 

import "strings"
import "github.com/GURURAJ8/banking/errors"

type NewAccountRequest struct {
	CustomerId  int     `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.CustomerId <= 0 {
		return errors.NewValidationError("CustomerId must be a positive integer")	
	}
	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidationError("AccountType must be either 'savings' or 'checking'")
	}
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a account	Minimum opening amount is 5000")
	}
	return nil
}