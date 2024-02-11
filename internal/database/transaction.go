package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Transaction struct {
	db        *sql.DB
	ID        string
	Amount    float64
	AccountId string
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

func (t *Transaction) Create(amount float64, accountId string) (Transaction, error) {
	id := uuid.New().String()

	_, err := t.db.Exec("INSERT INTO transactions (id, amount, account_id) VALUES ($1, $2, $3)",
		id, amount, accountId)

	if err != nil {
		return Transaction{}, err
	}

	return Transaction{ID: id, Amount: amount, AccountId: accountId}, nil
}

func (t *Transaction) FindAll() ([]Transaction, error) {
	rows, err := t.db.Query("SELECT id, amount, account_id FROM transactions")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := []Transaction{}

	for rows.Next() {
		var amount float64
		var id, accountId string

		if err := rows.Scan(&id, &amount, &accountId); err != nil {
			return nil, err
		}

		transactions = append(transactions, Transaction{ID: id, Amount: amount, AccountId: accountId})
	}

	return transactions, nil
}

func (t *Transaction) FindOneById(id string) (Transaction, error) {
	var amount float64
	var accountId string

	err := t.db.QueryRow("SELECT amount, account_id FROM transactions WHERE id = $1", id).
		Scan(&amount, &accountId)

	if err != nil {
		return Transaction{}, err
	}

	return Transaction{ID: id, Amount: amount, AccountId: accountId}, nil
}
