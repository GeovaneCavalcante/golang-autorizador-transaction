package transaction

import (
	"github.com/authorizer/entity"
	"github.com/authorizer/usecase/account"
)

//Service book usecase
type Service struct {
	accountService account.UseCase
	repo           Repository
}

//NewService create new service
func NewService(r Repository, accountService account.UseCase) *Service {
	return &Service{
		repo:           r,
		accountService: accountService,
	}
}

//CreateTransaction create transaction
func (s *Service) CreateTransaction(merchant string, amount int, timeData string) (*entity.Transaction, error) {
	violations := make([]string, 0, 0)

	t, err := entity.NewTransaction(merchant, amount, timeData, violations)
	if err != nil {
		return nil, err
	}

	t, valid, err := s.ValidateTransaction(t)
	if err != nil {
		return nil, entity.ErrValidateTransaction
	}
	if !valid {
		return t, nil
	}

	t, err = s.repo.Create(t)
	if err != nil {
		return nil, entity.ErrCreateTransaction
	}
	_, err = s.accountService.DeductAccountValue(amount)
	if err != nil {
		return nil, err
	}

	return t, nil
}

//ListTransactions getting all transactions
func (s *Service) ListTransactions() ([]*entity.Transaction, error) {
	c, err := s.repo.List()
	if err != nil {
		return nil, entity.ErrListTransactions
	}
	return c, nil
}

//ValidateTransaction validate transaction
func (s *Service) ValidateTransaction(t *entity.Transaction) (*entity.Transaction, bool, error) {
	violations := make([]string, 0, 0)
	valid := true

	accountVerify, err := s.accountService.AccountInitializationVerification()
	if err != nil {
		return nil, false, err
	}
	if !accountVerify {
		violations = append(violations, entity.AccountNotInitialized)
		t.Violations = violations
		return t, false, nil
	}

	activeCard, err := s.accountService.CheckActiveCard()
	if err != nil {
		return nil, false, err
	}
	if !activeCard {
		violations = append(violations, entity.CardNotActive)
		t.Violations = violations
		valid = false
	}

	enoughLimit, err := s.accountService.CheckBalanceForTransaction(t.Amount)
	if err != nil {
		return nil, false, err
	}
	if !enoughLimit {
		violations = append(violations, entity.InsufficientLimit)
		t.Violations = violations
		valid = false
	}

	authorization, err := s.CheckTransitionFrequency(t)
	if err != nil {
		return nil, false, err
	}
	if !authorization {
		violations = append(violations, entity.HighFrequencySmallInterval)
		t.Violations = violations
		valid = false
	}
	authorization, err = s.CheckDoubleTransaction(t)
	if !authorization {
		violations = append(violations, entity.DoubledTransaction)
		t.Violations = violations
		valid = false
	}

	t.Violations = violations

	return t, valid, nil
}

//CheckTransitionFrequency check transition frequency
func (s *Service) CheckTransitionFrequency(t *entity.Transaction) (bool, error) {
	transactions, err := s.ListTransactions()
	if err != nil {
		return false, err
	}
	valid := false
	if len(transactions) < 3 {
		return true, nil
	}
	for i := 1; i <= 3; i++ {
		transaction := transactions[len(transactions)-i]
		diffTime := t.Time.Sub(transaction.Time).Seconds()
		if diffTime > 120 {
			valid = true
		}
	}
	return valid, nil
}

//CheckDoubleTransaction check double transaction
func (s *Service) CheckDoubleTransaction(t *entity.Transaction) (bool, error) {
	transactions, err := s.ListTransactions()
	if err != nil {
		return false, err
	}
	valid := true
	if len(transactions) < 1 {
		return true, nil
	}
	for i := 1; i <= len(transactions); i++ {
		transaction := transactions[len(transactions)-i]
		diffTime := t.Time.Sub(transaction.Time).Seconds()
		if diffTime < 120 && transaction.Amount == t.Amount && transaction.Merchant == t.Merchant {
			valid = false
		}
	}
	return valid, nil
}
