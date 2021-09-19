package presenter

import (
	"github.com/authorizer/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccountWrapper(t *testing.T) {
	acc := entity.NewAccount(true, 100, make([]string, 0, 0))
	u := NewAccountWrapper(acc, make([]string, 0, 0))
	assert.Empty(t, u.Violations)
	assert.NotNil(t, u.Account)
}

func TestNewAccountWrapperEmpty(t *testing.T) {
	u := NewAccountWrapper(nil, make([]string, 0, 0))
	assert.Empty(t, u.Violations)
	assert.Equal(t, u.Account, make(map[string]interface{}))
}
