package repository

import (
	"github.com/authorizer/entity"
	"github.com/authorizer/usecase/transaction"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionInme_Create(t *testing.T) {
	m := NewTransactionInme()
	t1 := transaction.NewFixtureTransaction()
	ac, err := m.Create(&t1)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Transaction{}, ac)
	assert.NotNil(t, ac)
}

func TestTransactionInme_List(t *testing.T) {
	m := NewTransactionInme()
	t1 := transaction.NewFixtureTransaction()
	_, _ = m.Create(&t1)
	ac, err := m.List()
	assert.Nil(t, err)

	assert.NotNil(t, ac)
	assert.Equal(t, len(ac), 1)
}
