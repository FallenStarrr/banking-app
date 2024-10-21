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

func NewConn(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	c.Host, c.User, c.Password, c.Dbname, 
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		  return nil, err
	}
	return db, nil
}


type Repo struct {
	db  *gorm.DB
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