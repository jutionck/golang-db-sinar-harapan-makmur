package entity

import "github.com/google/uuid"

type Vehicle struct {
	Id             string
	Brand          string
	Model          string
	ProductionYear int `db:"production_year"`
	Color          string
	IsAutomatic    bool `db:"is_automatic"`
	Stock          int
	SalePrice      int    `db:"sale_price"`
	Status         string // enum: "Baru" & "Bekas"
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "Baru" || v.Status == "Bekas"
}

func (v *Vehicle) SetId() {
	v.Id = uuid.New().String()
}
