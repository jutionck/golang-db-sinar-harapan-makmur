package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type VehicleUseCase interface {
	RegisterNewVehicle(newVehicle entity.Vehicle) error
	FindAllVehicle() ([]entity.Vehicle, error)
	GetVehicle(id string) (entity.Vehicle, error)
	UpdateVehicle(newVehicle entity.Vehicle) error
	DeleteVehicle(id string) error
	Paging(requestQueryParams dto.RequestQueryParams) ([]entity.Vehicle, dto.Paging, error)
	Count(sql string) (int, error)
	GroupBy(selectedBy string, whereBy map[string]interface{}, groupBy string) ([]dto.VehicleGroupCountDto, error)
	UpdateVehicleStock(count int, id string) error
}

type vehicleUseCase struct {
	vehicleRepo repository.VehicleRepository
}

func (v *vehicleUseCase) RegisterNewVehicle(newVehicle entity.Vehicle) error {

	if err := vehicleValidation(newVehicle); err != nil {
		return err
	}

	err := v.vehicleRepo.Create(newVehicle)
	if err != nil {
		return fmt.Errorf("Failed to create new vehicle: %v", err)
	}

	return nil
}

func (v *vehicleUseCase) FindAllVehicle() ([]entity.Vehicle, error) {
	return v.vehicleRepo.List()
}

func (v *vehicleUseCase) GetVehicle(id string) (entity.Vehicle, error) {
	return v.vehicleRepo.Get(id)
}

func (v *vehicleUseCase) UpdateVehicle(newVehicle entity.Vehicle) error {
	if err := vehicleValidation(newVehicle); err != nil {
		return err
	}

	err := v.vehicleRepo.Update(newVehicle)
	if err != nil {
		return fmt.Errorf("Failed to udpate vehicle: %v", err)
	}

	return nil
}

func (v *vehicleUseCase) DeleteVehicle(id string) error {
	return v.vehicleRepo.Delete(id)
}

func (v *vehicleUseCase) Paging(requestQueryParams dto.RequestQueryParams) ([]entity.Vehicle, dto.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, dto.Paging{}, fmt.Errorf("Invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return v.vehicleRepo.Paging(requestQueryParams)
}

func vehicleValidation(payload entity.Vehicle) error {
	if payload.Brand == "" || payload.Model == "" || payload.Color == "" {
		return fmt.Errorf("Brand, Model, and Color are required fields")
	}

	if !payload.IsValidStatus() {
		return fmt.Errorf("Invalid status: %s", payload.Status)
	}

	if payload.SalePrice < 0 || payload.SalePrice == 0 {
		return fmt.Errorf("Sale price can't zero or negative ")
	}

	if payload.Stock < 0 {
		return fmt.Errorf("Stock can't negative ")
	}
	return nil
}

func (v *vehicleUseCase) Count(sql string) (int, error) {
	return v.vehicleRepo.Count(sql)
}

func (v *vehicleUseCase) GroupBy(selectedBy string, whereBy map[string]interface{}, groupBy string) ([]dto.VehicleGroupCountDto, error) {
	return v.vehicleRepo.GroupBy(selectedBy, whereBy, groupBy)
}

func (v *vehicleUseCase) UpdateVehicleStock(count int, id string) error {
	return v.vehicleRepo.UpdateStock(count, id)
}

func NewVehicleUseCase(vehicleRepo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{vehicleRepo: vehicleRepo}
}
