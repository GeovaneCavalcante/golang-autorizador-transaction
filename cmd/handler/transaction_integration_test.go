package handler

import (
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/entity"
	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) TestCreateTransaction() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), out.Violations, make([]string, 0, 0))
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationAccountNotInitialized() {
	input := `{"transaction": {"merchant": "Vivara", "amount": 100.0, "time": "2019-02-13T11:00:00.000Z"}}`
	val := NewInputData(input)
	out, err := MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.AccountNotInitialized)
	assert.Equal(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationCardNotActive() {
	input := `{"account": {"active-card": false, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.CardNotActive)
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationInsufficientLimit() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 2000, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.InsufficientLimit)
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationHighFrequencySmallInterval() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	input = `{"transaction": {"merchant": "Burger king", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	input = `{"transaction": {"merchant": "C&A", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	input = `{"transaction": {"merchant": "Ebay", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.HighFrequencySmallInterval)
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationDoubledTransaction() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	input = `{"transaction": {"merchant": "Vivara", "amount": 100, "time": "2019-02-13T11:00:00.000Z"}}`
	val = NewInputData(input)
	out, err = MakeTransactionHandlers(s.transactionService, s.accountService, val["transaction"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.DoubledTransaction)
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}
