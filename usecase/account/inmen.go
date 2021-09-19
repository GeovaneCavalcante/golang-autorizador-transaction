package account

import (
	"errors"
	"github.com/authorizer/entity"
)

// Repo in memory
type inmem struct {
	data      *entity.Account
	errGet    bool
	errCreate bool
	errUpdate bool
}

//newInmem create new repository in memory
func newInmem(errGet, errCreate, errUpdate bool) *inmem {
	return &inmem{
		errGet:    errGet,
		errCreate: errCreate,
		errUpdate: errUpdate,
	}
}

//Get account
func (r *inmem) Get() (*entity.Account, error) {
	if r.errGet {
		return nil, errors.New("Empty account")
	}
	return r.data, nil
}

//Create account
func (r *inmem) Create(e *entity.Account) (*entity.Account, error) {
	if r.errCreate {
		return nil, errors.New("Error create account")
	}
	r.data = e
	return e, nil
}

//Update account
func (r *inmem) Update(e *entity.Account) (*entity.Account, error) {
	if r.errUpdate {
		return nil, errors.New("Error update account")
	}
	r.data = e
	return e, nil
}
