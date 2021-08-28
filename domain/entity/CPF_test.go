package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidCPFOnlyZero(t *testing.T) {
	c := NewCPF("00000000000")
	assert.False(t, c.Valid())
}

func TestInvalidCPF(t *testing.T) {
	c := NewCPF("86446422799")
	assert.False(t, c.Valid())
}

func TestValidCPF(t *testing.T) {
	c := NewCPF("86446422784")
	assert.True(t, c.Valid())
}

func TestValidCPFWithSpecialCharacter(t *testing.T) {
	c := NewCPF("864.464.227-84")
	assert.True(t, c.Valid())
}

func TestValidCPFOther(t *testing.T) {
	c := NewCPF("91720489726")
	assert.True(t, c.Valid())
}

func TestInvalidCPFInvalidWithString(t *testing.T) {
	c := NewCPF("a1720489726")
	assert.False(t, c.Valid())
}
