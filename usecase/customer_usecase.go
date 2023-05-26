package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(newCustomer entity.Customer) error
	FindAllCustomer() ([]entity.Customer, error)
	GetCustomer(id string) (entity.Customer, error)
	UpdateCustomer(newCustomer entity.Customer) error
	DeleteCustomer(id string) error
	FindCustomerByEmail(email string) (entity.Customer, error)
	FindCustomerByPhoneNumber(phoneNumber string) (entity.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(newCustomer entity.Customer) error {
	isExists, _ := c.GetCustomer(newCustomer.Id)
	if isExists.Id == newCustomer.Id {
		return fmt.Errorf("Customer with ID: %v exists", newCustomer.Id)
	}

	isEmailExist, _ := c.customerRepo.GetByEmail(newCustomer.Email)
	if isEmailExist.Email == newCustomer.Email {
		return fmt.Errorf("Customer with email: %v exists", newCustomer.Email)
	}

	isPhoneNumberExist, _ := c.customerRepo.GetByPhoneNumber(newCustomer.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newCustomer.PhoneNumber {
		return fmt.Errorf("Customer with phone number: %v exists", newCustomer.PhoneNumber)
	}

	if newCustomer.FirstName == "" || newCustomer.LastName == "" || newCustomer.PhoneNumber == "" || newCustomer.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	err := c.customerRepo.Create(newCustomer)
	if err != nil {
		return fmt.Errorf("Failed to create new vehicle: %v", err)
	}

	return nil
}

func (c *customerUseCase) FindAllCustomer() ([]entity.Customer, error) {
	return c.customerRepo.List()
}

func (c *customerUseCase) GetCustomer(id string) (entity.Customer, error) {
	return c.customerRepo.Get(id)
}

func (c *customerUseCase) UpdateCustomer(newCustomer entity.Customer) error {
	isEmailExist, _ := c.customerRepo.GetByEmail(newCustomer.Email)

	if isEmailExist.Email == newCustomer.Email && isEmailExist.Id != newCustomer.Id {
		return fmt.Errorf("Customer with email: %v exists", newCustomer.Email)
	}

	isPhoneNumberExist, _ := c.customerRepo.GetByPhoneNumber(newCustomer.PhoneNumber)
	if isPhoneNumberExist.PhoneNumber == newCustomer.PhoneNumber && isPhoneNumberExist.Id != newCustomer.Id {
		return fmt.Errorf("Customer with phone number: %v exists", newCustomer.PhoneNumber)
	}

	if newCustomer.FirstName == "" || newCustomer.LastName == "" || newCustomer.PhoneNumber == "" || newCustomer.Email == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	err := c.customerRepo.Update(newCustomer)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.customerRepo.Delete(id)
}

func (c *customerUseCase) FindCustomerByEmail(email string) (entity.Customer, error) {
	return c.customerRepo.GetByEmail(email)
}
func (c *customerUseCase) FindCustomerByPhoneNumber(phoneNumber string) (entity.Customer, error) {
	return c.customerRepo.GetByPhoneNumber(phoneNumber)
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: customerRepo}
}
