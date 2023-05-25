package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
)

type EmployeeRepository interface {
	BaseRepository[entity.Employee]
	GetByEmail(email string) (entity.Employee, error)
	GetByPhoneNumber(phoneNumber string) (entity.Employee, error)
}

type employeeRepository struct {
	db *sqlx.DB
}

func (e *employeeRepository) Create(newData entity.Employee) error {
	sql := "INSERT INTO employee (id, first_name, last_name, address, phone_number, email, bod, position, salary, manager_id) VALUES (:id, :first_name, :last_name, :address, :phone_number, :email, :bod, :position, :salary, :manager_id)"
	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}

	namedArgs := map[string]interface{}{
		"id":           uuid.New().String(),
		"first_name":   newData.FirstName.String,
		"last_name":    newData.LastName.String,
		"address":      newData.Address.String,
		"phone_number": newData.PhoneNumber.String,
		"email":        newData.Email.String,
		"bod":          newData.Bod.Time,
		"position":     newData.Position.String,
		"salary":       newData.Salary.Int64,
		"manager_id":   managerID,
	}

	_, err := e.db.NamedExec(sql, namedArgs)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) List() ([]entity.Employee, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee`
	var employees []entity.Employee
	rows, err := e.db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee entity.Employee
		err := rows.StructScan(&employee)
		if err != nil {
			return nil, err
		}

		var manager entity.Employee
		sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
		err = e.db.Get(&manager, sqlManager, employee.Id)
		if err != nil {
			return nil, err
		}
		manager.Manager = nil // set manager menjadi nil agar tidak terjadi infinite loop dalam reference circular
		employee.Manager = &manager
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *employeeRepository) Get(id string) (entity.Employee, error) {
	sql := `SELECT id, first_name, last_name, address, phone_number, email, bod, position, salary FROM employee WHERE id = $1`
	var employee entity.Employee
	err := e.db.Get(&employee, sql, id)
	if err != nil {
		return entity.Employee{}, err
	}
	var manager entity.Employee
	sqlManager := `select m.id, m.first_name, m.last_name, m.address, m.email, m.phone_number, m.bod, m.salary, m.position from employee e left join employee m on m.id = e.manager_id where e.id = $1`
	err = e.db.Get(&manager, sqlManager, employee.Id)
	if err != nil {
		return entity.Employee{}, err
	}
	employee.Manager = &manager

	return employee, nil
}

func (e *employeeRepository) Update(newData entity.Employee) error {
	sql := "UPDATE employee SET first_name = :first_name, last_name = :last_name, address = :address, phone_number = :phone_number, email = :email, bod = :bod, position = :position, salary = :salary, manager_id = :manager_id WHERE id = :id"
	var managerID interface{}
	if newData.Manager != nil {
		managerID = newData.Manager.Id
	} else {
		managerID = nil
	}
	namedArgs := map[string]interface{}{
		"id":           newData.Id.String,
		"first_name":   newData.FirstName.String,
		"last_name":    newData.LastName.String,
		"address":      newData.Address.String,
		"phone_number": newData.PhoneNumber.String,
		"email":        newData.Email.String,
		"bod":          newData.Bod.Time,
		"position":     newData.Position.String,
		"salary":       newData.Salary.Int64,
		"manager_id":   managerID,
	}

	_, err := e.db.NamedExec(sql, namedArgs)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	sql := "DELETE FROM employee WHERE id = $1"
	_, err := e.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) GetByEmail(email string) (entity.Employee, error) {
	sql := `SELECT id, email FROM employee WHERE email = $1`
	var employee entity.Employee
	err := e.db.Get(&employee, sql, email)
	if err != nil {
		return entity.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetByPhoneNumber(phoneNumber string) (entity.Employee, error) {
	sql := `SELECT id, phone_number FROM employee WHERE phone_number = $1`
	var employee entity.Employee
	err := e.db.Get(&employee, sql, phoneNumber)
	if err != nil {
		return entity.Employee{}, err
	}
	return employee, nil
}

func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
