package repository

import "github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"

type BaseRepository[T any] interface {
	Create(newData T) error
	List() ([]T, error)
	Get(id string) (T, error)
	Update(newData T) error
	Delete(id string) error
}

type BaseRepositoryPaging[T any] interface {
	Paging(requestQueryParams dto.RequestQueryParams) ([]T, dto.Paging, error)
}

type BaseRepositoryAggregate[T any] interface {
	Count(sql string) (int, error)
	GroupBy(selectedBy string, whereBy map[string]interface{}, groupBy string) ([]T, error)
}
