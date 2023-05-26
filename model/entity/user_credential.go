package entity

type UserCredential struct {
	Id       string `json:"id"`
	UserName string `db:"user_name" json:"username"`
	Password string `json:"password"`
	IsActive bool   `db:"is_active" json:"isActive"`
}
