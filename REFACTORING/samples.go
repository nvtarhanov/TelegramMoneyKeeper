package REFACTORING

import (
	"errors"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
)

type Repository interface {
	GetData(id int) model.Account
	SetData(*model.Account) error
}

//FIRST REPO

type ReposytoryDatabase struct{}

func (r ReposytoryDatabase) GetData(id int) model.Account {

	return model.Account{}
}

func (r ReposytoryDatabase) SetData(*model.Account) error {

	return errors.New("asd")
}

///////////////////////SECOND REPO
type MockRepo struct{}

func (r MockRepo) GetData(id int) model.Account {

	return model.Account{}
}

func (r MockRepo) SetData(*model.Account) error {

	return errors.New("asd")
}

//HOW TO USE
func UseThisStuff() {

	// r := ReposytoryDatabase{}

	// // m := r.GetData(1)

	// var r Repository

	// r = ReposytoryDatabase{}

	// r.
}
