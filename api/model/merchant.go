package model

type Merchant struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (Merchant) TableName() string {
	return "merchants"
}
