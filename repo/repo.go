package repo

import (
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
	"github.com/FallenStarrr/banking-app/config"
	"github.com/FallenStarrr/banking-app/model"
	"context"
	"fmt"
)


type AccRepo interface {
	    GetAcc(id string) model.AccountModel
			CreateAcc(acc model.AccountModel)
			UpdateAcc(id string)
			DeleteAcc(id string)
			BlockAcc(id string)
}


type Repo struct {
	db  *
}


func (r Repo) GetAcc(id string) model.AccountModel {

}


func (r Repo) CreateAcc(acc model.AccountModel)  {
	
}


func (r Repo) UpdateAcc(id string)  {
	var ctx = context.Background()

}


func (r Repo) DeleteAcc(id string)  {
	var ctx = context.Background()

}

func (r Repo) BlockAcc(id string)  {
	var ctx = context.Background()

}

func NewRepo()  AccRepo {

}