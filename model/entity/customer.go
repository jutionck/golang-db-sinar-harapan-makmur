package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          string
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Address     string
	PhoneNumber string `db:"phone_number"`
	Email       string
	Bod         time.Time
}

func (c *Customer) SetId() {
	c.Id = uuid.New().String()
}
