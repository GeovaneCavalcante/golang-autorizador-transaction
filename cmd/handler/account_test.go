// +build !integration

package handler

import (
	"errors"
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/usecase/account"
	accountMock "github.com/authorizer/usecase/account/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeAccountHandlers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	accountService := accountMock.NewMockUseCase(controller)
	ac := account.NewFixtureAccount()
	accountService.EXPECT().
		CreateAccount(gomock.Any(), gomock.Any()).
		Return(&ac, nil)

	acc := map[string]interface{}{
		"active-card":     true,
		"available-limit": 1000.0,
	}
	out, err := MakeAccountHandlers(accountService, acc)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.IsType(t, &presenter.AccountWrapper{}, out)
}

func TestMakeTransactionHandlersError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ac := account.NewFixtureAccount()
	accountService := accountMock.NewMockUseCase(controller)
	accountService.EXPECT().
		CreateAccount(gomock.Any(), gomock.Any()).
		Return(&ac, errors.New("Error create account"))

	acc := map[string]interface{}{
		"active-card":     true,
		"available-limit": 1000.0,
	}
	out, err := MakeAccountHandlers(accountService, acc)
	assert.NotNil(t, err)
	assert.Nil(t, out)
}
