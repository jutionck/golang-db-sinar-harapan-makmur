package entity

import "github.com/google/uuid"

type Vehicle struct {
	Id             string `json:"id"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	ProductionYear int    `db:"production_year" json:"productionYear"`
	Color          string `json:"color"`
	IsAutomatic    bool   `db:"is_automatic" json:"isAutomatic"`
	Stock          int    `json:"stock"`
	SalePrice      int    `db:"sale_price" json:"salePrice"`
	Status         string `json:"status"` // enum: "Baru" & "Bekas"
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "Baru" || v.Status == "Bekas"
}

func (v *Vehicle) SetId() {
	v.Id = uuid.New().String()
}
