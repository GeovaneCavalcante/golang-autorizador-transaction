package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNewTransaction(t *testing.T) {
	u, err := NewTransaction("Burger King", 10, "2019-02-13T11:01:31.000Z", make([]string, 0, 0))
	assert.Nil(t, err)
	assert.Equal(t, u.Merchant, "Burger King")
	assert.NotNil(t, u.ID)
	assert.Equal(t, u.Amount, 10)
}

func TestNewNewTransactionErrorTimeFormat(t *testing.T) {
	u, err := NewTransaction("Burger King", 10, "teste", make([]string, 0, 0))
	assert.NotNil(t, err)
	assert.Nil(t, u)
}
