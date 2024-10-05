package model

type Customer struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Balance  float64 `json:"balance"`
}

func (Customer) TableName() string {
	return "customers"
}
