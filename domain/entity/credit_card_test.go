package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "José da Silva", 12, 2024, 123)
	assert.Equal(t, nil, err)

	_, err = NewCreditCard("123", "José da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())
}
