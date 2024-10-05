package model

type PaymentRequest struct {
	MerchantID uint    `json:"merchant_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}
