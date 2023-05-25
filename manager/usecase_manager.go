package manager

import "github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"

type UseCaseManager interface {
	VehicleUseCase() usecase.VehicleUseCase
	CustomerUseCase() usecase.CustomerUseCase
	EmployeeUseCase() usecase.EmployeeUseCase
	TransactionUseCase() usecase.TransactionUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repoManager.TransactionRepo(), u.VehicleUseCase(), u.CustomerUseCase(), u.EmployeeUseCase())
}

func (u *useCaseManager) VehicleUseCase() usecase.VehicleUseCase {
	return usecase.NewVehicleUseCase(u.repoManager.VehicleRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
