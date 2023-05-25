package dto

import (
	"time"
)

type EmployeeDto struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Bod         time.Time `json:"bod"`
	Position    string    `json:"position"`
	Salary      int64     `json:"salary"`
	ManagerId   string    `json:"managerId"`
}
