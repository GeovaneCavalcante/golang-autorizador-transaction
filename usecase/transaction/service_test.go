package transaction

import (
	"errors"
	"github.com/authorizer/entity"
	"github.com/authorizer/usecase/account"
	"github.com/authorizer/usecase/account/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestCreateTransactionErrorNewEntityTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(false, errors.New("Error authorization")).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrValidateTransaction)
	assert.Nil(t, tr)
}

func TestCreateTransactionErrorInvalid(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(false, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	assert.Nil(t, err)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestCreateTransactionError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(true, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrCreateTransaction)
	assert.Nil(t, tr)
}
func TestCreateTransactionErrorDeductAccountValue(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, errors.New("Error deduct value")).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrCreateTransaction)
	assert.Nil(t, tr)
}

func TestCreateTransactionErrorValidateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, err := m.CreateTransaction(t1.Merchant, t1.Amount, "teste")
	assert.NotNil(t, err)
	assert.Nil(t, tr)
}

func TestValidateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(false, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, false)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestValidateTransactionFailAccountVerify(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, true)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestValidateTransactionFailActiveCard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(false, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, false)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestValidateTransactionErrorActiveCard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(false, errors.New("Error check active card")).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, valid, false)
}

func TestValidateTransactionFailEnoughLimit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(false, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, false)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestValidateTransactionErrorEnoughLimit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(false, errors.New("Error check balance for transaction")).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, valid, false)
}

func TestValidateTransactionFailCheckTransitionFrequency(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction("Riachuelo", t1.Amount, "2019-02-13T11:02:00.000Z")
	_, _ = m.CreateTransaction("Ebay", t1.Amount, "2019-02-13T11:03:00.000Z")
	_, _ = m.CreateTransaction("Burger max", t1.Amount, "2019-02-13T11:05:00.000Z")
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, false)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestValidateTransactionErrorCheckTransitionFrequency(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, true)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.NotNil(t, valid, true)
}

func TestValidateTransactionFailCheckDoubleTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	tr, valid, err := m.ValidateTransaction(&t1)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, valid, false)
	assert.IsType(t, &entity.Transaction{}, tr)
}

func TestListTransactions(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction("Riachuelo", t1.Amount, "2019-02-13T11:02:00.000Z")
	tr, err := m.ListTransactions()
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, len(tr), 2)
}

func TestListTransactionsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	repo := newInmem(false, true)
	m := NewService(repo, service)
	tr, err := m.ListTransactions()
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrListTransactions)
	assert.Nil(t, tr)
}

func TestCheckTransitionFrequency(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction("Riachuelo", t1.Amount, "2019-02-13T11:02:00.000Z")
	_, _ = m.CreateTransaction("Ebay", t1.Amount, "2019-02-13T11:03:00.000Z")
	_, _ = m.CreateTransaction("Burger max", t1.Amount, "2019-02-13T11:05:00.000Z")
	t2 := NewFixtureTransaction()
	ac, err := m.CheckTransitionFrequency(&t2)
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, false)
}
func TestCheckTransitionFrequencyErrorLisTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	repo := newInmem(false, true)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	ac, err := m.CheckTransitionFrequency(&t1)
	assert.NotNil(t, err)
	assert.Equal(t, ac, false)
}

func TestCheckDoubleTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		AccountInitializationVerification().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckActiveCard().
		Return(true, nil).AnyTimes()
	service.EXPECT().
		CheckBalanceForTransaction(gomock.Any()).
		Return(true, nil).AnyTimes()
	acc := account.NewFixtureAccount()
	service.EXPECT().
		DeductAccountValue(gomock.Any()).
		Return(&acc, nil).AnyTimes()
	repo := newInmem(false, false)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:00:00.000Z")
	_, _ = m.CreateTransaction(t1.Merchant, t1.Amount, "2019-02-13T11:02:00.000Z")
	t2 := NewFixtureTransaction()
	ac, err := m.CheckDoubleTransaction(&t2)
	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ac, false)
}

func TestCheckDoubleTransactionErrorLisTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	repo := newInmem(false, true)
	m := NewService(repo, service)
	t1 := NewFixtureTransaction()
	ac, err := m.CheckDoubleTransaction(&t1)
	assert.NotNil(t, err)
	assert.Equal(t, ac, false)
}
