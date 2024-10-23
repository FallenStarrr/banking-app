package model

type TransactionErr struct {
	TransactionId string `json:"transaction_id"`
	From int `json:"from"`
	To int `json:"to"`
	Amount float64 `json:"amount"`
	Message string `json:"message"`
	StatusCode int `json:status_code`
}