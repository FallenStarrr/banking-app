package controller

import (
	"github.com/gofiber/fiber/v2"
"github.com/FallenStarrr/banking-app/config"
"github.com/FallenStarrr/banking-app/model"
"github.com/FallenStarrr/banking-app/srv"
"github.com/FallenStarrr/banking-app/repo"
)


   var c  = config.GetConfig()
	 

	var db,  _ = repo.NewConn(&c)
 

	var r  = repo.NewRepo(db)
	var s  = srv.NewAccountSrv(r)



func GetAccount(c *fiber.Ctx) error {
	return c.JSON(s.GetAcc(c.Params("id")))
}


func DeleteAccount(c *fiber.Ctx) error {
	return s.DeleteAcc(c.Params("id"))
}


func PutAccount(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


func CreateAccount(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}