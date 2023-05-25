package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type CustomerRepository interface {
	BaseRepository[entity.Customer]
	GetByEmail(email string) (entity.Customer, error)
	GetByPhoneNumber(phoneNumber string) (entity.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Create(newData entity.Customer) error {
	sql := "INSERT INTO customer (id, first_name, last_name, address, phone_number, email, bod) VALUES (:id, :first_name, :last_name, :address, :phone_number, :email, :bod)"
	_, err := c.db.NamedExec(sql, &newData)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) List() ([]entity.Customer, error) {
	var customers []entity.Customer
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod FROM customer`
	err := c.db.Select(&customers, sql)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) Get(id string) (entity.Customer, error) {
	var customer entity.Customer
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod FROM customer WHERE id = $1`
	err := c.db.Get(&customer, sql, id)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) Update(newData entity.Customer) error {
	sql := "UPDATE customer SET first_name = :first_name, last_name = :last_name, address = :address, phone_number = :phone_number, email = :email, bod = :bod WHERE id = :id"
	_, err := c.db.NamedExec(sql, &newData)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id string) error {
	sql := "DELETE FROM customer WHERE id = $1"
	_, err := c.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) GetByEmail(email string) (entity.Customer, error) {
	var customer entity.Customer
	sql := `SELECT id, email FROM customer WHERE email = $1`
	err := c.db.Get(&customer, sql, email)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) GetByPhoneNumber(phoneNumber string) (entity.Customer, error) {
	var customer entity.Customer
	sql := `SELECT id, phone_number FROM customer WHERE phone_number = $1`
	err := c.db.Get(&customer, sql, phoneNumber)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db: db}
}
