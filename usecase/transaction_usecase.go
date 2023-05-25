package usecase

import (
	"fmt"
	"time"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type TransactionUseCase interface {
	RegisterNewTransaction(newData entity.Transaction) error
	FindAllTransaction() ([]dto.TransactionResponseDto, error)
	FindTransactionById(id string) (dto.TransactionResponseDto, error)
	GetAll() ([]entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	vehicleUseCase  VehicleUseCase
	customerUseCase CustomerUseCase
	employeeUseCase EmployeeUseCase
}

func (t *transactionUsecase) RegisterNewTransaction(newData entity.Transaction) error {
	// get vehicle
	vehicle, err := t.vehicleUseCase.GetVehicle(newData.Vehicle.Id)
	if err != nil {
		return fmt.Errorf("vehicle with ID: %s not exists", newData.Vehicle.Id)
	}

	if vehicle.Stock < newData.Qty {
		return fmt.Errorf("stock of vehicle is not enough")
	}

	// get customer
	customer, err := t.customerUseCase.GetCustomer(newData.Customer.Id)
	if err != nil {
		return fmt.Errorf("customer with ID: %s not exists", newData.Customer.Id)
	}

	// get employee
	employee, err := t.employeeUseCase.GetEmployee(newData.Employee.Id.String)
	if err != nil {
		return fmt.Errorf("employee with ID: %v not exists", newData.Employee.Id)
	}

	newData.Vehicle = vehicle
	newData.Customer = customer
	newData.Employee = employee
	newData.TransactionDate = time.Now()
	newData.PaymentAmount = vehicle.SalePrice

	err = t.vehicleUseCase.UpdateVehicleStock(newData.Qty, vehicle.Id)
	if err != nil {
		return fmt.Errorf("failed update vehicle stock")
	}

	return t.transactionRepo.Create(newData)
}

func (t *transactionUsecase) FindAllTransaction() ([]dto.TransactionResponseDto, error) {
	return t.transactionRepo.List()
}

func (t *transactionUsecase) GetAll() ([]entity.Transaction, error) {
	return t.transactionRepo.GetAll()
}

func (t *transactionUsecase) FindTransactionById(id string) (dto.TransactionResponseDto, error) {
	return t.transactionRepo.Get(id)
}

func NewTransactionUseCase(
	transactionRepo repository.TransactionRepository,
	vehicleUseCase VehicleUseCase,
	customerUseCase CustomerUseCase,
	employeeUseCase EmployeeUseCase) TransactionUseCase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		vehicleUseCase:  vehicleUseCase,
		customerUseCase: customerUseCase,
		employeeUseCase: employeeUseCase,
	}
}
