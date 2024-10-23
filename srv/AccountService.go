package srv

import (
	"github.com/FallenStarrr/banking-app/model"
	"github.com/FallenStarrr/banking-app/repo"
	"gorm.io/gorm"
	"fmt"
)

type AccSrv interface {
	GetAcc(id string) *model.AccountModel
	CreateAcc(acc *model.AccountModel)  (tx *gorm.DB)
	UpdateAcc(id string, field string, value string)  (tx *gorm.DB, err string)
	DeleteAcc(id string) (tx *gorm.DB)
	SendMoney(from, to int, amount float64)  error
	BlockAcc(id string) string 
}


type AccountService struct {
	 repo repo.AccRepo
	}

	func NewAccountSrv(repo repo.AccRepo) AccSrv {
			return &AccountService{repo: repo}
	}





func(a *AccountService) GetAcc(id string) *model.AccountModel {
			return a.repo.GetAcc(id)
}


func(a *AccountService) DeleteAcc(id string)  (tx *gorm.DB)  {
			return  a.repo.DeleteAcc(id)
}


func(a *AccountService) CreateAcc( acc *model.AccountModel)  (tx *gorm.DB)  {
	 return a.repo.CreateAcc(acc)
}


func(a *AccountService) UpdateAcc(id string, field string, value string)  (tx *gorm.DB, err string) {
	tx, err  = a.repo.UpdateAcc(id, field, value)
	fmt.Println(err)
	return
}

func (a *AccountService) SendMoney(from, to int, amount float64)  error {
	return a.repo.SendMoney(from, to, amount)
}

func  (a *AccountService) BlockAcc(id string) string {
			return a.repo.BlockAcc(id)
} 

