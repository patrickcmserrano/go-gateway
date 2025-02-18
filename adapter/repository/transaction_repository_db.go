package repository

import (
	"database/sql"
	"time"
)

// TransactionRepositoryDB is a struct to hold the database connection
type TransactionRepositoryDB struct {
	DB *sql.DB
}

// NewTransactionRepositoryDB is a function to create a new TransactionRepositoryDB
func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{DB: db}
}

// Close is a method to close the database connection
func (r *TransactionRepositoryDB) Close() {
	r.DB.Close()
}

func (t *TransactionRepositoryDB) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err1 := t.DB.Prepare(`
		INSERT INTO transactions (id, account, amount, status, error_message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err1 != nil {
		return err1
	}

	_, err2 := stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())

	if err2 != nil {
		return err2
	}

	return nil
}
