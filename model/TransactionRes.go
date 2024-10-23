package model


type TransactionRes struct {
	TransactionId string `json:"transaction_id"`
	From int `json:"from"`
	To int `json:"to"`
	Amount float64 `json:"amount"`
}