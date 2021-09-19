package transaction

import "github.com/authorizer/entity"

//Reader interface
type Reader interface {
	List() ([]*entity.Transaction, error)
}

//Writer account writer
type Writer interface {
	Create(e *entity.Transaction) (*entity.Transaction, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateTransaction(merchant string, amount int, timeData string) (*entity.Transaction, error)
	ListTransactions() ([]*entity.Transaction, error)
	ValidateTransaction(t *entity.Transaction) (*entity.Transaction, bool, error)
	CheckTransitionFrequency(t *entity.Transaction) (bool, error)
	CheckDoubleTransaction(t *entity.Transaction) (bool, error)
}
