package process_transaction

import (
	"github.com/patrickcmserrano/go-gateway/domain/entity"
	"github.com/patrickcmserrano/go-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	// Tenta criar um novo cartão de crédito
	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if invalidCC != nil {
		// Se o cartão de crédito for inválido, insere uma transação rejeitada no repositório
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		// Retorna a saída da transação com status REJECTED e a mensagem de erro
		return TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}, nil
	}

	// Associa o cartão de crédito à transação
	transaction.SetCreditCard(*cc)
	// Verifica se a transação é válida
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		// Se a transação for inválida, insere uma transação rejeitada no repositório
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidTransaction.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}
		// Retorna a saída da transação com status REJECTED e a mensagem de erro
		return TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidTransaction.Error(),
		}, nil
	}

	// Se a transação for válida, insere uma transação aprovada no repositório
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	// Retorna a saída da transação com status APPROVED
	return TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}, nil
}
