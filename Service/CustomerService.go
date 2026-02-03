package Service	
//This file contains the service layer for Customer
// primary port of customer service layer
import "github.com/GURURAJ8/banking/domain"
import "github.com/GURURAJ8/banking/errors"
import "github.com/GURURAJ8/banking/dto"

type CustomerService interface{
	GetAllCustomers() ([]dto.CustomerResponse, *errors.AppError)
	GetCustomerById(id int) (*dto.CustomerResponse, *errors.AppError)
}		

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errors.AppError){
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.NewNotFoundError("Failed to retrieve customers")
	}
	var responses []dto.CustomerResponse
	for _, c := range customers {
		responses = append(responses, c.ToDto())
	}
	return responses, nil
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{r}
}

func (s DefaultCustomerService) GetCustomerById(id int) (*dto.CustomerResponse, *errors.AppError){
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response:= c.ToDto()
	return &response, nil
}