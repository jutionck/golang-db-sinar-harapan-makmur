package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type TransactionRepository interface {
	Create(newData entity.Transaction) error
	List() ([]dto.TransactionResponseDto, error)
	GetAll() ([]entity.Transaction, error)
	Get(id string) (dto.TransactionResponseDto, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func (t *transactionRepository) Create(newData entity.Transaction) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}
	sql := "INSERT INTO transaction (id, transaction_date, vehicle_id, customer_id, employee_id, type, qty, payment_amount) VALUES (:id, :transaction_date, :vehicle_id, :customer_id, :employee_id, :type, :qty, :payment_amount)"

	namedArgs := map[string]interface{}{
		"id":               newData.Id,
		"transaction_date": newData.TransactionDate,
		"vehicle_id":       newData.Vehicle.Id,
		"customer_id":      newData.Customer.Id,
		"employee_id":      newData.Employee.Id,
		"type":             newData.Type,
		"qty":              newData.Qty,
		"payment_amount":   newData.PaymentAmount,
	}

	result, err := tx.NamedExec(sql, namedArgs)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	} else if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no rows affected")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *transactionRepository) List() ([]dto.TransactionResponseDto, error) {
	sql := `
	select 
	t.id,
	t.transaction_date, 
	c.id as customer_id,
	c.first_name || ' ' || c.last_name as customer_name,
	v.id as vehicle_id,
	v.brand,
	v.model,
	e.id as employee_id,
	e.first_name || ' ' || e.last_name as employee_name,
	t.qty,
	t.type,
	t.payment_amount
from
	transaction t
inner join vehicle v on v.id = t.vehicle_id 
inner join customer c on c.id = t.customer_id 
inner join employee e on e.id = t.employee_id`

	var transactions []dto.TransactionResponseDto
	err := t.db.Select(&transactions, sql)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) Get(id string) (dto.TransactionResponseDto, error) {
	sql := `
	select 
	t.id,
	t.transaction_date, 
	c.id as customer_id,
	c.first_name || ' ' || c.last_name as customer_name,
	v.id as vehicle_id,
	v.brand,
	v.model,
	e.id as employee_id,
	e.first_name || ' ' || e.last_name as employee_name,
	t.qty,
	t.type,
	t.payment_amount
from
	transaction t
inner join vehicle v on v.id = t.vehicle_id 
inner join customer c on c.id = t.customer_id 
inner join employee e on e.id = t.employee_id
where t.id = $1`

	var transaction dto.TransactionResponseDto
	err := t.db.Get(&transaction, sql, id)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}
	return transaction, nil
}

func (t *transactionRepository) GetAll() ([]entity.Transaction, error) {
	sql := "SELECT id,transaction_date,customer_id,vehicle_id,employee_id, qty, type, payment_amount FROM transaction ORDER BY transaction_date ASC"

	var transactions []entity.Transaction
	rows, err := t.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.TransactionDate,
			&transaction.Customer.Id,
			&transaction.Vehicle.Id,
			&transaction.Employee.Id,
			&transaction.Qty,
			&transaction.Type,
			&transaction.PaymentAmount,
		)
		if err != nil {
			return nil, err
		}
		sqlVehicle := "SELECT id, brand, model, production_year, color, is_automatic, sale_price, stock, status FROM vehicle WHERE id = $1"
		var vehicle entity.Vehicle
		err = t.db.Get(&vehicle, sqlVehicle, transaction.Vehicle.Id)
		if err != nil {
			return nil, err
		}
		transaction.Vehicle = vehicle

		sqlCustomer := "SELECT id, first_name, last_name, address, email, phone_number, bod FROM customer WHERE id = $1"
		var customer entity.Customer
		err = t.db.Get(&customer, sqlCustomer, transaction.Customer.Id)
		if err != nil {
			return nil, err
		}
		transaction.Customer = customer

		sqlEmployee := "SELECT id, first_name, last_name, address, email, phone_number, bod, salary, position FROM employee WHERE id = $1"
		var employee entity.Employee
		err = t.db.Get(&employee, sqlEmployee, transaction.Employee.Id)
		if err != nil {
			return nil, err
		}

		sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
		var manager entity.Employee
		err = t.db.Get(&manager, sqlManager, employee.Id)
		if err != nil {
			return nil, err
		}
		manager.Manager = nil // set manager menjadi nil agar tidak terjadi infinite loop dalam reference circular
		employee.Manager = &manager
		transaction.Employee = employee
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
