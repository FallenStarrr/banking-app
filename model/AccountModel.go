package model

import "gorm.io/gorm"

type AccountModel struct {
	 gorm.Model
	 Type string `json:"type"`
	 Status string `json:"status"`
	 Balance float64 `json:"balance"`
	 ClientName string	`json:"client_name"`
	 Iin string  `json:"iin"`
}