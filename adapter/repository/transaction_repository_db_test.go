package repository

import (
	"os"
	"testing"

	"github.com/patrickcmserrano/go-gateway/adapter/repository/fixture"
	"github.com/patrickcmserrano/go-gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repo := NewTransactionRepositoryDB(db)
	err := repo.Insert("1", "1", 1.0, entity.APPROVED, "")
	assert.Nil(t, err)
}
