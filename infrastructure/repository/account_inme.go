package repository

import (
	"github.com/authorizer/entity"
)

//AccountInme in memory repo
type AccountInme struct {
	data *entity.Account
}

//NewAccountInme create new repository
func NewAccountInme() *AccountInme {
	return &AccountInme{}
}

//Get account
func (r *AccountInme) Get() (*entity.Account, error) {
	return r.data, nil
}

//Create account
func (r *AccountInme) Create(e *entity.Account) (*entity.Account, error) {
	r.data = e
	return e, nil
}

//Update account
func (r *AccountInme) Update(e *entity.Account) (*entity.Account, error) {
	r.data = e
	return e, nil
}
