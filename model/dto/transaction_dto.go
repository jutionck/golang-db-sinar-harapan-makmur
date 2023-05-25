package dto

import "time"

type TransactionResponseDto struct {
	Id              string
	TransactionDate time.Time `db:"transaction_date"`
	CustomerId      string    `db:"customer_id"`
	CustomerName    string    `db:"customer_name"`
	VehicleId       string    `db:"vehicle_id"`
	VehicleBrand    string    `db:"brand"`
	VehicleModel    string    `db:"model"`
	EmployeeId      string    `db:"employee_id"`
	EmployeeName    string    `db:"employee_name"`
	Qty             int
	Type            string
	PaymentAmount   int `db:"payment_amount"`
}
