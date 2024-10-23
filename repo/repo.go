package repo

import (
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
	"github.com/FallenStarrr/banking-app/config"
	"github.com/FallenStarrr/banking-app/model"
	_"context"
	"fmt"
	"errors"
)


type AccRepo interface {
	    GetAcc(id string) *model.AccountModel
			CreateAcc(acc *model.AccountModel) (tx *gorm.DB)
			UpdateAcc(id string, field string, value string) (tx *gorm.DB)
			DeleteAcc(id string) (tx *gorm.DB)
			BlockAcc(id string)
			SendMoney(from, to int, amount float64)  error
}

func NewConn(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",c.Host, c.User, c.Password, c.Db, c.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		  return nil, err
	}


	return db, nil
}


type Repo struct {
	Db  *gorm.DB
}


func (r Repo) GetAcc(id string) *model.AccountModel {
     var acc model.AccountModel 
	   r.Db.First(&acc, id)
		 return &acc
}


func (r Repo) CreateAcc(acc *model.AccountModel) (tx *gorm.DB)  {
	
	return r.Db.Create(&model.AccountModel{
		Type: acc.Type, 
		Status: acc.Status,
		Balance: 0,
		ClientName: acc.ClientName,
		Iin: acc.Iin,
	})
}


func (r Repo) UpdateAcc(id string, field string, value string) (tx *gorm.DB)  {
	var acc model.AccountModel
	r.Db.First(&acc, id)
	return r.Db.Model(&acc).Update(field,  value)
}


func (r Repo) DeleteAcc(id string) (tx *gorm.DB)  {
	var acc model.AccountModel
	return r.Db.Delete(&acc, id)
}

func (r Repo) BlockAcc(id string)  {
	var acc model.AccountModel
	r.Db.First(&acc, id)
	r.Db.Model(&acc).Update("status",  "blocked")
}

func (r Repo) SendMoney(from, to int, amount float64) error {

	var fromAcc model.AccountModel
	var toAcc model.AccountModel  
	r.Db.First(&fromAcc, from)
	r.Db.First(&toAcc, to)
	if fromAcc.Status == "blocked" &&  toAcc.Status == "blocked" {
		return errors.New("Blocked user!")
	}
	if fromAcc.Balance <= 0 {
			return errors.New("Non enough money")
	} 
	return r.Db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
	fromAmount := fromAcc.Balance - amount
  toAmount := toAcc.Balance + amount
		if err := tx.Model(&fromAcc).Update("balance",  fromAmount).Error; err != nil {
			// return any error will rollback
			return err
		}
	
		if err := tx.Model(&toAcc).Update("balance",  toAmount).Error; err != nil {
			// return any error will rollback
			return err
		}
	
		// return nil will commit the whole transaction
		return nil
	})
}
func NewRepo(db  *gorm.DB)  AccRepo {
				return Repo{Db: db}
}