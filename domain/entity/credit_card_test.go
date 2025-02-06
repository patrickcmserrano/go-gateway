package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "José da Silva", 12, 2030, 123)
	assert.Equal(t, nil, err)

	_, err = NewCreditCard("123", "José da Silva", 12, 2030, 123)
	assert.Equal(t, "invalid credit card number", err.Error())
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "José da Silva", 13, 2030, 123)
	assert.Equal(t, "invalid credit card expiration month", err.Error())
}

func TestCreditCardExpirationYear(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "José da Silva", 12, 2022, 123)
	assert.Equal(t, "invalid credit card expiration year", err.Error())
}
