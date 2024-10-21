package main

import "github.com/FallenStarrr/banking-app/config"
import "fmt"

func main() {
   c := config.GetConfig()
	fmt.Println(c)
}