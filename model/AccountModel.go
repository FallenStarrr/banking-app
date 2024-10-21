package model

type AccountModel struct {
	 Type string `redis:"type"`
	 Id int `redis:"id"`
	 RegistrationDate string `redis:"registration_date"`
	 Balance float64 `redis:"balance"`
	 ClientName string	`redis:"client_name"`
	 Iin string  `redis:"iin"`
}