/*// +build integration
 */
package handler

import (
	"encoding/json"
	"github.com/authorizer/infrastructure/repository"
	"github.com/authorizer/usecase/account"
	"github.com/authorizer/usecase/transaction"
	"github.com/stretchr/testify/suite"
	"testing"
)

type integrationTestSuite struct {
	suite.Suite
	accountService     *account.Service
	transactionService *transaction.Service
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &integrationTestSuite{})
}

func (s *integrationTestSuite) SetupTest() {
	accountRepo := repository.NewAccountInme()
	s.accountService = account.NewService(accountRepo)
	transactionRepo := repository.NewTransactionInme()
	s.transactionService = transaction.NewService(transactionRepo, s.accountService)
}

func (s *integrationTestSuite) TearDownTest() {
	s.accountService = nil
	s.transactionService = nil
}

func NewInputData(data string) map[string]map[string]interface{} {
	var input map[string]map[string]interface{}
	json.Unmarshal([]byte(data), &input)
	return input
}
