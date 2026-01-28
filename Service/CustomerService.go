package Service	
//This file contains the service layer for Customer
import "github.com/GURURAJ8/banking/domain"

type CustomerService interface{
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id int) (*domain.Customer, errors.AppError)
}		

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{r}
}

func (s DefaultCustomerService) GetCustomerById(id int) (*domain.Customer, errors.AppError){
	return s.repo.ById(id)
}