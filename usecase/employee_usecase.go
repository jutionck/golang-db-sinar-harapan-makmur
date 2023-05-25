package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(newEmployee entity.Employee) error
	FindAllEmployee() ([]entity.Employee, error)
	GetEmployee(id string) (entity.Employee, error)
	UpdateEmployee(newEmployee entity.Employee) error
	DeleteEmployee(id string) error
	FindEmployeeByEmail(email string) (entity.Employee, error)
	FindEmployeeByPhoneNumber(phoneNumber string) (entity.Employee, error)
	FindManagerById(id string) (entity.Employee, error)
}

type employeeUseCase struct {
	employeeRepo repository.EmployeeRepository
}

func (e *employeeUseCase) RegisterNewEmployee(newEmployee entity.Employee) error {
	isExists, _ := e.GetEmployee(newEmployee.Id.String)
	if isExists.Id == newEmployee.Id {
		return fmt.Errorf("Employee with ID: %v exists", newEmployee.Id)
	}

	isEmailExist, _ := e.employeeRepo.GetByEmail(newEmployee.Email.String)
	if isEmailExist.Email == newEmployee.Email {
		return fmt.Errorf("Employee with email: %v exists", newEmployee.Email)
	}

	isPhoneNumberExist, _ := e.employeeRepo.GetByPhoneNumber(newEmployee.PhoneNumber.String)
	if isPhoneNumberExist.PhoneNumber == newEmployee.PhoneNumber {
		return fmt.Errorf("Employee with phone number: %v exists", newEmployee.PhoneNumber)
	}

	if newEmployee.FirstName.String == "" || newEmployee.LastName.String == "" || newEmployee.PhoneNumber.String == "" || newEmployee.Email.String == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	if newEmployee.Manager != nil {
		manager, err := e.FindManagerById(newEmployee.Manager.Id.String)
		if err != nil {
			return fmt.Errorf("Manager with ID: %v exists", newEmployee.Manager.Id)
		}
		newEmployee.Manager = &manager
	} else {
		newEmployee.Manager = nil
	}

	err := e.employeeRepo.Create(newEmployee)
	if err != nil {
		return fmt.Errorf("Failed to create new employee: %v", err)
	}

	return nil
}

func (e *employeeUseCase) FindAllEmployee() ([]entity.Employee, error) {
	return e.employeeRepo.List()
}

func (e *employeeUseCase) GetEmployee(id string) (entity.Employee, error) {
	employee, err := e.employeeRepo.Get(id)
	if err != nil {
		return entity.Employee{}, fmt.Errorf("Employee with ID: %v not found", id)
	}
	return employee, nil
}

func (e *employeeUseCase) UpdateEmployee(newEmployee entity.Employee) error {
	isEmailExist, _ := e.employeeRepo.GetByEmail(newEmployee.Email.String)

	if isEmailExist.Email == newEmployee.Email && isEmailExist.Id != newEmployee.Id {
		return fmt.Errorf("Employee with email: %v exists", newEmployee.Email)
	}

	isPhoneNumberExist, _ := e.employeeRepo.GetByPhoneNumber(newEmployee.PhoneNumber.String)
	if isPhoneNumberExist.PhoneNumber == newEmployee.PhoneNumber && isPhoneNumberExist.Id != newEmployee.Id {
		return fmt.Errorf("Employee with phone number: %v exists", newEmployee.PhoneNumber)
	}

	if newEmployee.FirstName.String == "" || newEmployee.LastName.String == "" || newEmployee.PhoneNumber.String == "" || newEmployee.Email.String == "" {
		return fmt.Errorf("FirstName, LastName, PhoneNumber and Email are required fields")
	}

	if newEmployee.Manager != nil {
		manager, err := e.FindManagerById(newEmployee.Manager.Id.String)
		if err != nil {
			return fmt.Errorf("Manager with ID: %v exists", newEmployee.Manager.Id)
		}
		newEmployee.Manager = &manager
	} else {
		newEmployee.Manager = nil
	}

	err := e.employeeRepo.Update(newEmployee)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	return e.employeeRepo.Delete(id)
}

func (e *employeeUseCase) FindEmployeeByEmail(email string) (entity.Employee, error) {
	return e.employeeRepo.GetByEmail(email)
}

func (e *employeeUseCase) FindEmployeeByPhoneNumber(phoneNumber string) (entity.Employee, error) {
	return e.employeeRepo.GetByPhoneNumber(phoneNumber)
}

func (e *employeeUseCase) FindManagerById(id string) (entity.Employee, error) {
	return e.GetEmployee(id)
}

func NewEmployeeUseCase(employeeRepo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{employeeRepo: employeeRepo}
}
