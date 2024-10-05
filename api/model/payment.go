package model

type Payment struct {
	ID          uint    `json:"id"`
	CustomerID  uint    `json:"customer_id"`
	MerchantID  uint    `json:"merchant_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Transaction string  `json:"transaction"`
}

func (Payment) TableName() string {
	return "payments"
}
