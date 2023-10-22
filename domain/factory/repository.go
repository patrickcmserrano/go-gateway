package factory

import "github.com/patrickcmserrano/go-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
