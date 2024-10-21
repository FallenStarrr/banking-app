package main

import (
	"github.com/FallenStarrr/banking-app/config"
 "github.com/FallenStarrr/banking-app/repo"
 "github.com/FallenStarrr/banking-app/model"
 "fmt"
)

func main() {
   c := config.GetConfig()
	 fmt.Println(c)

	db, err := repo.NewConn(c)
	if err != nil {
		fmt.Println(err)
	}
  r := repo.NewRepo(db)
	r.CreateAcc(&model.AccountModel{Type: "Savings", Status: "Active", ClientName: "John Doe", Id: "1", Iin: "249" })
	r.CreateAcc(&model.AccountModel{Type: "Savings", Status: "Active", ClientName: "Jane Doe", Id: "2", Iin: "250" })
}