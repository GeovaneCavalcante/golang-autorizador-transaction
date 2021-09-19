package handler

import (
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/usecase/account"
)

//MakeAccountHandlers make account handlers
func MakeAccountHandlers(service account.UseCase, accountData map[string]interface{}) (*presenter.AccountWrapper, error) {
	activeCard := accountData["active-card"].(bool)
	availableLimit := int(accountData["available-limit"].(float64))
	acc, err := service.CreateAccount(activeCard, availableLimit)
	if err != nil {
		return nil, err
	}
	accountPresenter := presenter.NewAccountWrapper(acc, acc.Violations)
	return accountPresenter, nil
}
