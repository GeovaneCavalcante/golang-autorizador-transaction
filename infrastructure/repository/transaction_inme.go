package repository

import (
	"github.com/authorizer/entity"
)

//TransactionInme in memory repo
type TransactionInme struct {
	data []*entity.Transaction
}

//NewTransactionInme create new repository
func NewTransactionInme() *TransactionInme {
	return &TransactionInme{}
}

//Create transaction
func (r *TransactionInme) Create(e *entity.Transaction) (*entity.Transaction, error) {
	r.data = append(r.data, e)
	return e, nil
}

//List transaction
func (r *TransactionInme) List() ([]*entity.Transaction, error) {
	return r.data, nil
}
