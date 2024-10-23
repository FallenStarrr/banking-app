package srv

import (
	"github.com/FallenStarrr/banking-app/model"
	"github.com/FallenStarrr/banking-app/repo"
)

type AccSrv interface {
	GetAcc(id int) *model.AccountModel
	CreateAcc(acc *model.AccountModel)
	UpdateAcc(id int, field string, value string)
	DeleteAcc(id int)
}


type AccountService struct {
	 repo repo.AccRepo
	}

	func NewAccountSrv(repo repo.AccRepo) AccSrv {
			return &AccountService{repo: repo}
	}





func(a *AccountService) GetAcc(id int) *model.AccountModel {
			return a.repo.GetAcc(id)
}


func(a *AccountService) DeleteAcc(id int)  {
			 a.repo.DeleteAcc(id)
}


func(a *AccountService) CreateAcc( acc *model.AccountModel)  {
	 a.repo.CreateAcc(acc)
}


func(a *AccountService) UpdateAcc(id int, field string, value string) {
	a.repo.UpdateAcc(id, field, value)
}

