package main

import (
	"github.com/FallenStarrr/banking-app/config"
 "github.com/FallenStarrr/banking-app/repo"
 "github.com/FallenStarrr/banking-app/model"
 "fmt"
)

func main() {
   c := config.GetConfig()
	 fmt.Println(c.Db)

	db, err := repo.NewConn(&c)
	if err != nil {
		fmt.Println("DB ERR : ", err)
		return
	}
  r := repo.NewRepo(db)
	// r.CreateAcc(&model.AccountModel{Type: "Savings", Status: "Active", ClientName: "John Doe", Iin: "249" })
	// r.CreateAcc(&model.AccountModel{Type: "Savings", Status: "Active", ClientName: "Jane Doe", Iin: "250" })
	r.DeleteAcc(3)
	r.DeleteAcc(4)
	r.DeleteAcc(5)
	r.DeleteAcc(6)
	r.SendMoney(1, 2, 1000)


}