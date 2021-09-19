package handler

import (
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/usecase/account"
	"github.com/authorizer/usecase/transaction"
)

//MakeTransactionHandlers make transaction handlers
func MakeTransactionHandlers(transactionService transaction.UseCase, accountService account.UseCase, transactionData map[string]interface{}) (*presenter.AccountWrapper, error) {
	merchant := transactionData["merchant"].(string)
	amount := int(transactionData["amount"].(float64))
	time := transactionData["time"].(string)

	t, err := transactionService.CreateTransaction(merchant, amount, time)
	if err != nil {
		return nil, err
	}
	acc, err := accountService.GetAccount()
	if err != nil {
		return nil, err
	}
	accountPresenter := presenter.NewAccountWrapper(acc, t.Violations)
	return accountPresenter, nil
}
