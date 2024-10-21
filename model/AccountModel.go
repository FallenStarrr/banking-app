package model

type AccountModel struct {
	 Type string `status:"type"`
	 Status string `json:"status"`
	 Id string `status:"id"`
	 RegistrationDate string `status:"registration_date"`
	 Balance float64 `status:"balance"`
	 ClientName string	`status:"client_name"`
	 Iin string  `status:"iin"`
}