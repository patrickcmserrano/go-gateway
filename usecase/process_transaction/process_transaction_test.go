package process_transaction

import (
	"github.com/patrickcmserrano/go-gateway/domain/entity"
	mock_repository "github.com/patrickcmserrano/go-gateway/domain/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "400000000000000000000000",
		CreditCardName:            "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectedCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    12002,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "Credit Card does not have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
