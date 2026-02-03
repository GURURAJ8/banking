package Service

//primary port of account service layer
import "github.com/GURURAJ8/banking/domain"
import "github.com/GURURAJ8/banking/errors"
import "github.com/GURURAJ8/banking/dto"
import "time"

type AccountService interface{
	NewAccountService(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
}

type DefaultAccountService struct{
	repo domain.AccountRepository
}

func (s DefaultAccountService)NewAccountService(r dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}
	domainAccount := domain.Account{
		CustomerId:  r.CustomerId,
		AccountType: r.AccountType,
		Amount:      r.Amount,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"), // This would typically be the current date
		Status: "1",
	}

	newAccount, appErr := s.repo.Save(domainAccount)
	if appErr != nil {
		return nil, appErr
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(r domain.AccountRepository) DefaultAccountService{
	return DefaultAccountService{r}
}