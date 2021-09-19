package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	ID := NewID()
	assert.NotEmpty(t, ID)
}
