package account

import "github.com/authorizer/entity"

//Reader interface
type Reader interface {
	Get() (*entity.Account, error)
}

//Writer account writer
type Writer interface {
	Create(e *entity.Account) (*entity.Account, error)
	Update(e *entity.Account) (*entity.Account, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetAccount() (*entity.Account, error)
	CreateAccount(activeCard bool, availableLimit int) (*entity.Account, error)
	DeductAccountValue(amount int) (*entity.Account, error)
	AccountInitializationVerification() (bool, error)
	CheckActiveCard() (bool, error)
	CheckBalanceForTransaction(amount int) (bool, error)
}
