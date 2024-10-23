package main

import (
	"github.com/FallenStarrr/banking-app/controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/account/:id", controller.GetAccount)

	app.Delete("/account/:id", controller.DeleteAccount)
	app.Post("/account", controller.CreateAccount)

	app.Put("/account/:id", controller.PutAccount)
	app.Put("/block/:id", controller.BlockAccount)
	app.Put("/send/account", controller.SendMoney)

	log.Fatal(app.Listen(":3000"))

}
