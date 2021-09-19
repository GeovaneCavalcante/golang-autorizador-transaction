package transaction

import (
	"errors"
	"github.com/authorizer/entity"
)

// Repo in memory
type inmem struct {
	data      []*entity.Transaction
	errCreate bool
	errList   bool
}

//newInmem create new repository in memory
func newInmem(errCreate, errList bool) *inmem {
	return &inmem{
		errCreate: errCreate,
		errList:   errList,
	}
}

//Create transaction
func (r *inmem) Create(e *entity.Transaction) (*entity.Transaction, error) {
	if r.errCreate {
		return nil, errors.New("Error create transaction")
	}
	r.data = append(r.data, e)
	return e, nil
}

//List transaction
func (r *inmem) List() ([]*entity.Transaction, error) {
	if r.errList {
		return nil, errors.New("Error list transactions")
	}
	return r.data, nil
}
