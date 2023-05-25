package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type Employee struct {
	Id          sql.NullString
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	Address     sql.NullString
	PhoneNumber sql.NullString `db:"phone_number"`
	Email       sql.NullString
	Bod         sql.NullTime
	Position    sql.NullString
	Salary      sql.NullInt64
	Manager     *Employee `db:"manager_id"`
}

func (e *Employee) SetId() {
	e.Id = sql.NullString{String: uuid.New().String()}
}
