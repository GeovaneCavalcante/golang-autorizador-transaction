package account

import (
	"github.com/authorizer/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateAccount create account
func (s *Service) CreateAccount(activeCard bool, availableLimit int) (*entity.Account, error) {
	violations := make([]string, 0, 0)
	c, err := s.GetAccount()
	if err != nil {
		return nil, entity.ErrGetAccount
	}
	if c != nil {
		violations = append(violations, entity.AccountAlreadyInitialized)
		c.Violations = violations
		return c, nil
	}
	c = entity.NewAccount(activeCard, availableLimit, violations)
	c, err = s.repo.Create(c)
	if err != nil {
		return nil, entity.ErrCreateAccount
	}
	return c, nil
}

//DeductAccountValue update account limit
func (s *Service) DeductAccountValue(amount int) (*entity.Account, error) {
	c, err := s.GetAccount()
	if err != nil {
		return nil, entity.ErrGetAccount
	}
	c.AvailableLimit -= amount
	c, err = s.repo.Update(c)
	if err != nil {
		return nil, entity.ErrUpdateAccount
	}
	return c, err
}

//GetAccount get account
func (s *Service) GetAccount() (*entity.Account, error) {
	c, err := s.repo.Get()
	if err != nil {
		return nil, entity.ErrGetAccount
	}
	return c, nil
}

//AccountInitializationVerification account Initialization Verification
func (s *Service) AccountInitializationVerification() (bool, error) {
	c, err := s.GetAccount()
	if err != nil {
		return false, entity.ErrGetAccount
	}
	if c == nil {
		return false, nil
	}
	return true, nil
}

//CheckActiveCard check if the card is active
func (s *Service) CheckActiveCard() (bool, error) {
	c, err := s.GetAccount()
	if err != nil {
		return false, entity.ErrGetAccount
	}
	return c.ActiveCard, nil
}

//CheckBalanceForTransaction check if you have available balance for a transaction
func (s *Service) CheckBalanceForTransaction(amount int) (bool, error) {
	c, err := s.GetAccount()
	if err != nil {
		return false, entity.ErrGetAccount
	}
	if c.AvailableLimit >= amount {
		return true, nil
	}
	return false, nil
}
