package repository

import (
	"github.com/authorizer/entity"
	"github.com/authorizer/usecase/account"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountInme_Create(t *testing.T) {
	m := NewAccountInme()
	c1 := account.NewFixtureAccount()
	ac, err := m.Create(&c1)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotNil(t, ac)
}

func TestAccountInme_Get(t *testing.T) {
	m := NewAccountInme()
	c1 := account.NewFixtureAccount()
	_, _ = m.Create(&c1)
	ac, err := m.Get()
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotNil(t, ac)
}

func TestAccountInme_Update(t *testing.T) {
	m := NewAccountInme()
	c1 := account.NewFixtureAccount()
	c2 := account.NewFixtureAccount()
	_, _ = m.Create(&c1)
	ac, err := m.Update(&c2)
	assert.Nil(t, err)
	assert.IsType(t, &entity.Account{}, ac)
	assert.NotNil(t, ac)
	assert.Equal(t, ac.AvailableLimit, c2.AvailableLimit)
}
