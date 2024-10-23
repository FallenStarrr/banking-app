package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FallenStarrr/banking-app/controller"
	"log"
)

func main() {

	 
	app := fiber.New()

	app.Get("/account/:id", controller.GetAccount)

	app.Delete("/account/:id", controller.DeleteAccount)
	app.Post("/account", controller.CreateAccount)

	app.Put("/account/:id", controller.PutAccount)

	log.Fatal(app.Listen(":3000"))

}
