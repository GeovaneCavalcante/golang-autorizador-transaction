// +build !integration

package handler

import (
	"errors"
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/usecase/account"
	accountMock "github.com/authorizer/usecase/account/mock"
	transactionUseCase "github.com/authorizer/usecase/transaction"
	transactionMock "github.com/authorizer/usecase/transaction/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeTransactionHandlers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	transactionService := transactionMock.NewMockUseCase(controller)
	tr := transactionUseCase.NewFixtureTransaction()
	transactionService.EXPECT().
		CreateTransaction(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&tr, nil)
	accountService := accountMock.NewMockUseCase(controller)
	ac := account.NewFixtureAccount()
	accountService.EXPECT().
		GetAccount().
		Return(&ac, nil)

	transaction := map[string]interface{}{
		"merchant": "Vivara",
		"amount":   1250.0,
		"time":     "2019-02-13T11:00:00.000Z",
	}
	out, err := MakeTransactionHandlers(transactionService, accountService, transaction)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.IsType(t, &presenter.AccountWrapper{}, out)
}

func TestMakeTransactionHandlersErrorCreateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	transactionService := transactionMock.NewMockUseCase(controller)
	tr := transactionUseCase.NewFixtureTransaction()
	transactionService.EXPECT().
		CreateTransaction(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&tr, errors.New("Error create transaction"))
	accountService := accountMock.NewMockUseCase(controller)
	transaction := map[string]interface{}{
		"merchant": "Vivara",
		"amount":   1250.0,
		"time":     "2019-02-13T11:00:00.000Z",
	}
	out, err := MakeTransactionHandlers(transactionService, accountService, transaction)
	assert.NotNil(t, err)
	assert.Nil(t, out)
}

func TestMakeTransactionHandlersErrorGetAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	transactionService := transactionMock.NewMockUseCase(controller)
	tr := transactionUseCase.NewFixtureTransaction()
	transactionService.EXPECT().
		CreateTransaction(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&tr, nil)
	accountService := accountMock.NewMockUseCase(controller)
	ac := account.NewFixtureAccount()
	accountService.EXPECT().
		GetAccount().
		Return(&ac, errors.New("Error get transaction"))

	transaction := map[string]interface{}{
		"merchant": "Vivara",
		"amount":   1250.0,
		"time":     "2019-02-13T11:00:00.000Z",
	}
	out, err := MakeTransactionHandlers(transactionService, accountService, transaction)
	assert.NotNil(t, err)
	assert.Nil(t, out)
}
