package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          string    `json:"id"`
	FirstName   string    `db:"first_name" json:"firstName"`
	LastName    string    `db:"last_name" json:"lastName"`
	Address     string    `json:"address"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber"`
	Email       string    `json:"email"`
	Bod         time.Time `json:"bod"`
}

func (c *Customer) SetId() {
	c.Id = uuid.New().String()
}
