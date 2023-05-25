package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              string    `json:"id"`
	TransactionDate time.Time `db:"transaction_date" json:"transactionDate"`
	Vehicle         Vehicle   `db:"vehicle_id" json:"vehicle"`
	Customer        Customer  `db:"customer_id" json:"customer"`
	Employee        Employee  `db:"employee_id" json:"employee"`
	Type            string    `json:"type"` // enum: "Online" & "Offline"
	Qty             int       `json:"qty"`
	PaymentAmount   int       `db:"payment_amount" json:"paymentAmount"`
}

func (t *Transaction) IsValidType() bool {
	return t.Type == "Online" || t.Type == "Offline"
}

func (t *Transaction) SetId() {
	t.Id = uuid.New().String()
}
