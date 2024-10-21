package main

import (
	"github.com/FallenStarrr/banking-app/config"
 "github.com/FallenStarrr/banking-app/repo"
 "fmt"
)

func main() {
   c := config.GetConfig()
	 fmt.Println(c)
	 
	db, err := repo.NewConn(c)
	if err != nil {
		fmt.Println(err)
	}
	
}