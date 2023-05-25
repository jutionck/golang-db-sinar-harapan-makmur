package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/utils/common"
)

type VehicleRepository interface {
	BaseRepository[entity.Vehicle]
	BaseRepositoryPaging[entity.Vehicle]
	BaseRepositoryAggregate[dto.VehicleGroupCountDto]
	UpdateStock(count int, id string) error
}

type vehicleRepository struct {
	db *sqlx.DB
}

func (v *vehicleRepository) Create(newData entity.Vehicle) error {
	sql := "INSERT INTO vehicle (id, brand, model, production_year, color, is_automatic, sale_price, stock, status) VALUES (:id, :brand, :model, :production_year, :color, :is_automatic, :sale_price, :stock, :status)"
	_, err := v.db.NamedExec(sql, newData)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) List() ([]entity.Vehicle, error) {
	var vehicles []entity.Vehicle
	sql := `SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle`
	err := v.db.Select(&vehicles, sql)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (v *vehicleRepository) Get(id string) (entity.Vehicle, error) {
	var vehicle entity.Vehicle
	sql := `SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle WHERE id = $1`
	err := v.db.Get(&vehicle, sql, id)
	if err != nil {
		return entity.Vehicle{}, err
	}
	return vehicle, nil
}

func (v *vehicleRepository) Update(newData entity.Vehicle) error {
	sql := "UPDATE vehicle set brand = :brand, model = :model, production_year = :production_year, color = :color, is_automatic = :is_automatic, sale_price = :sale_price, stock = :stock, status = :status WHERE id = :id"
	_, err := v.db.NamedExec(sql, &newData)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) Delete(id string) error {
	sql := "DELETE FROM vehicle WHERE id = $1"
	_, err := v.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehicleRepository) Paging(requestQueryParams dto.RequestQueryParams) ([]entity.Vehicle, dto.Paging, error) {
	var paginationQuery dto.PaginationQuery
	var vehicles []entity.Vehicle
	paginationQuery = common.GetPaginationParams(requestQueryParams.PaginationParam)
	orderQuery := "ORDER BY id"
	if requestQueryParams.QueryParams.Order != "" && requestQueryParams.QueryParams.Sort != "" {
		sorting := "ASC"
		if requestQueryParams.QueryParams.Sort == "desc" {
			sorting = "DESC"
		}
		orderQuery = fmt.Sprintf("ORDER BY %s %s", requestQueryParams.QueryParams.Order, sorting)
	}
	sql := fmt.Sprintf("SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle %s LIMIT $1 OFFSET $2", orderQuery)
	err := v.db.Select(&vehicles, sql, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	totalRows, err := v.Count("SELECT COUNT(*) FROM vehicle")
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return vehicles, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (v *vehicleRepository) Count(sql string) (int, error) {
	row := v.db.QueryRowx(sql)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (v *vehicleRepository) GroupBy(selectedBy string, whereBy map[string]interface{}, groupBy string) ([]dto.VehicleGroupCountDto, error) {
	var vehicles []dto.VehicleGroupCountDto

	// Build the SQL query
	query := fmt.Sprintf("SELECT %s, COUNT(*) AS total_count FROM vehicle", selectedBy)
	if len(whereBy) > 0 {
		query += " WHERE "
		for k, v := range whereBy {
			query += fmt.Sprintf("%s=%v AND ", k, v)
		}
		query = strings.TrimSuffix(query, " AND ")
	}
	query += fmt.Sprintf(" GROUP BY %s", groupBy)

	// Execute the query
	rows, err := v.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map the query result to entity.Vehicle objects
	for rows.Next() {
		var vehicle dto.VehicleGroupCountDto
		err := rows.Scan(&vehicle.FieldName, &vehicle.FieldCount)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (v *vehicleRepository) UpdateStock(count int, id string) error {
	sql := "UPDATE vehicle SET stock = stock - $1 WHERE id = $2"
	_, err := v.db.Exec(sql, count, id)
	if err != nil {
		return err
	}
	return nil
}

func NewVehicleRepository(db *sqlx.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}
