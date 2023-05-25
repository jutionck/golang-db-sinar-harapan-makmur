package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              string
	TransactionDate time.Time `db:"transaction_date"`
	Vehicle         Vehicle   `db:"vehicle_id"`
	Customer        Customer  `db:"customer_id"`
	Employee        Employee  `db:"employee_id"`
	Type            string    // enum: "Online" & "Offline"
	Qty             int
	PaymentAmount   int `db:"payment_amount"`
}

func (t *Transaction) IsValidType() bool {
	return t.Type == "Online" || t.Type == "Offline"
}

func (t *Transaction) SetId() {
	t.Id = uuid.New().String()
}
