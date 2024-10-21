package repo

import (
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
	"github.com/FallenStarrr/banking-app/config"
	"github.com/FallenStarrr/banking-app/model"
	"context"
	"fmt"
	"errors"
)


type AccRepo interface {
	    GetAcc(id string) *model.AccountModel
			CreateAcc(acc *model.AccountModel)
			UpdateAcc(id string, field string, value string)
			DeleteAcc(id string)
			BlockAcc(id string)
			SendMoney(from, to string, amount float64)  error
}

func NewConn(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	c.Host, c.User, c.Password, c.Dbname, 
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		  return nil, err
	}

	db.AutoMigrate(&model.AccountModel{})
	return db, nil
}


type Repo struct {
	db  *gorm.DB
}


func (r Repo) GetAcc(id string) *model.AccountModel {
     var acc model.AccountModel 
	   r.db.First(&acc, id)
		 r.db.AutoMigrate(&model.AccountModel{})
		 return &acc
}


func (r Repo) CreateAcc(acc *model.AccountModel)  {
	
	r.db.Create(&model.AccountModel{
		Type: acc.Type, 
		Status: acc.Status,
		Id: acc.Id,
		Balance: 0,
		ClientName: acc.ClientName,
		Iin: acc.Iin,
	})
	r.db.AutoMigrate(&model.AccountModel{})
}


func (r Repo) UpdateAcc(id string, field string, value string)  {
	var acc model.AccountModel
	r.db.First(&acc, id)
	r.db.Model(&acc).Update(field,  value)
	r.db.AutoMigrate(&model.AccountModel{})
}


func (r Repo) DeleteAcc(id string)  {
	var acc model.AccountModel
	r.db.Delete(&acc, id)
	r.db.AutoMigrate(&model.AccountModel{})
}

func (r Repo) BlockAcc(id string)  {
	var acc model.AccountModel
	r.db.First(&acc, id)
	r.db.Model(&acc).Update("status",  "blocked")
	r.db.AutoMigrate(&model.AccountModel{})
}

func (r Repo) SendMoney(from, to string, amount float64) error {

	var fromAcc model.AccountModel
	var toAcc model.AccountModel  
	r.db.First(&fromAcc, from)
	r.db.First(&toAcc, to)

	if fromAcc.Balance <= 0 {
			return errors.New("Non enough money")
	} 
	return r.db.Transaction(func(tx *gorm.DB) error {
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
				return Repo{db: db}
}