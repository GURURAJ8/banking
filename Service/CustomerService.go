package Service	
//This file contains the service layer for Customer
import "github.com/GURURAJ8/banking/domain"

type CustomerService struct{
GetAllCustomers() ([]domain.Customer, error)}		

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{
		repo: r,
	}
}