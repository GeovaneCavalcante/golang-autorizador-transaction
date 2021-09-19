package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	u := NewAccount(true, 100, make([]string, 0, 0))
	assert.Equal(t, u.ActiveCard, true)
	assert.NotNil(t, u.ID)
	assert.Equal(t, u.AvailableLimit, 100)
}
