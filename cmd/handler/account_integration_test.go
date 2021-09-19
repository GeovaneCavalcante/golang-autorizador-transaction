package handler

import (
	"github.com/authorizer/cmd/presenter"
	"github.com/authorizer/entity"
	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) TestCreateAccount() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), out.Violations, make([]string, 0, 0))
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}

func (s *e2eTestSuite) TestCreateAccountViolationAccountAlreadyInitialized() {
	input := `{"account": {"active-card": true, "available-limit": 1000}}`
	val := NewInputData(input)
	out, err := MakeAccountHandlers(s.accountService, val["account"])
	out, err = MakeAccountHandlers(s.accountService, val["account"])
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), out)
	assert.IsType(s.T(), &presenter.AccountWrapper{}, out)
	assert.Equal(s.T(), len(out.Violations), 1)
	assert.Equal(s.T(), out.Violations[0], entity.AccountAlreadyInitialized)
	assert.NotEqual(s.T(), out.Account, make(map[string]interface{}))
}
