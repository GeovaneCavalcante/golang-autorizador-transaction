package account

import (
	"github.com/authorizer/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	ac, err := m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotNil(t, ac)
}

func TestCreateAccountErrorAccountAlreadyInitialized(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	c2 := NewFixtureAccount()
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.CreateAccount(c2.ActiveCard, c2.AvailableLimit)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotEmpty(t, ac.Violations)
	assert.NotNil(t, ac)
}

func TestCreateAccountErrorGetAccount(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	c2 := NewFixtureAccount()
	ac, err := m.CreateAccount(c2.ActiveCard, c2.AvailableLimit)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrGetAccount)
	assert.Nil(t, ac)
}

func TestCreateAccountError(t *testing.T) {
	repo := newInmem(false, true, false)
	m := NewService(repo)
	c2 := NewFixtureAccount()
	ac, err := m.CreateAccount(c2.ActiveCard, c2.AvailableLimit)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrCreateAccount)
	assert.Nil(t, ac)
}

func TestUpdateAccountLimit(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	c1.AvailableLimit = 1000
	ac, err := m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err = m.DeductAccountValue(200)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.Equal(t, ac.AvailableLimit, 800)
	assert.NotNil(t, ac)
}

func TestUpdateAccountLimitErrorGetAccount(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	c2 := NewFixtureAccount()
	ac, err := m.CreateAccount(c2.ActiveCard, c2.AvailableLimit)
	ac, err = m.DeductAccountValue(200)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrGetAccount)
	assert.Nil(t, ac)
}

func TestUpdateAccountLimitError(t *testing.T) {
	repo := newInmem(false, false, true)
	m := NewService(repo)
	c2 := NewFixtureAccount()
	ac, err := m.CreateAccount(c2.ActiveCard, c2.AvailableLimit)
	ac, err = m.DeductAccountValue(200)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrUpdateAccount)
	assert.Nil(t, ac)
}

func TestGetAccount(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.GetAccount()
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotNil(t, ac)
}

func TestGetAccountEmpty(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	ac, err := m.GetAccount()
	assert.NotNil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.Nil(t, ac)
}

func TestAccountInitializationVerification(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.AccountInitializationVerification()
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, true)
}

func TestAccountInitializationVerificationErrorGetAccount(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	ac, err := m.AccountInitializationVerification()
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrGetAccount)
	assert.Equal(t, ac, false)
}

func TestAccountInitializationVerificationAccountEmpty(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	ac, err := m.AccountInitializationVerification()
	assert.Nil(t, err)
	assert.Equal(t, ac, false)
}

func TestCheckActiveCard(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.CheckActiveCard()
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, true)
}

func TestCheckActiveCardErrorGetAccount(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	ac, err := m.CheckActiveCard()
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrGetAccount)
	assert.Equal(t, ac, false)
}

func TestCheckCheckBalanceForTransaction(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.CheckBalanceForTransaction(100)
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, true)
}

func TestCheckCheckBalanceForTransactionWithoutbalance(t *testing.T) {
	repo := newInmem(false, false, false)
	m := NewService(repo)
	c1 := NewFixtureAccount()
	c1.AvailableLimit = 100
	_, _ = m.CreateAccount(c1.ActiveCard, c1.AvailableLimit)
	ac, err := m.CheckBalanceForTransaction(200)
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, false)
}

func TestCheckCheckBalanceForTransactionErrorGetAccount(t *testing.T) {
	repo := newInmem(true, false, false)
	m := NewService(repo)
	ac, err := m.CheckBalanceForTransaction(100)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrGetAccount)
	assert.Equal(t, ac, false)
}
