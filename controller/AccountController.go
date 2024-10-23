package controller

import (
	"fmt"
	"strconv"
	"github.com/google/uuid"
	"github.com/FallenStarrr/banking-app/config"
	"github.com/FallenStarrr/banking-app/model"
	"github.com/FallenStarrr/banking-app/repo"
	"github.com/FallenStarrr/banking-app/srv"
	"github.com/gofiber/fiber/v2"
)


   var c  = config.GetConfig()
	 

	var db,  _ = repo.NewConn(&c)
 

	var r  = repo.NewRepo(db)
	var s  = srv.NewAccountSrv(r)



func GetAccount(c *fiber.Ctx) error {
	return c.JSON(s.GetAcc(c.Params("id")))
}


func DeleteAccount(c *fiber.Ctx) error {
	id :=  c.Params("id")
	s.DeleteAcc(id)
	return c.JSON(id)
}


func PutAccount(c *fiber.Ctx) error {
	id :=  c.Params("id")
	field :=  c.Query("field")
	val := c.Query("value")
	s.UpdateAcc(id, field, val)
	return c.JSON(id)
}


func SendMoney(c *fiber.Ctx) error {
	
	from :=  c.Query("from")
	
	f, err := strconv.Atoi(from)
	if err != nil {
		fmt.Println("FROM : ", err )
	}
	to :=  c.Query("to")
	t, _ := strconv.Atoi(to)
	amount := c.Query("amount")
	amt, _ := strconv.ParseFloat(amount, 64)
	s.SendMoney(f, t, amt)
	transactionId := uuid.New().String()
	r := model.TransactionRes{From: f, To: t, Amount: amt, TransactionId: transactionId}
	return c.JSON(r)
}


func CreateAccount(c *fiber.Ctx) error {

	var acc = &model.AccountModel{}
	if err := c.BodyParser(acc); err != nil {
		return err
}

 
	return c.JSON("Account created!")
}