package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Account struct {
	db    *sql.DB
	ID    string
	Name  string
	Email string
}

func NewAccount(db *sql.DB) *Account {
	return &Account{db: db}
}

func (a *Account) Create(name string, email string) (Account, error) {
	id := uuid.New().String()

	_, err := a.db.Exec("INSERT INTO accounts (id, name, email) VALUES ($1, $2, $3)",
		id, name, email)

	if err != nil {
		return Account{}, err
	}

	return Account{ID: id, Name: name, Email: email}, nil
}

func (a *Account) FindAll() ([]Account, error) {
	rows, err := a.db.Query("SELECT id, name, email FROM accounts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []Account{}

	for rows.Next() {
		var id, name, email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		accounts = append(accounts, Account{ID: id, Name: name, Email: email})
	}

	return accounts, nil
}

func (a *Account) FindOneById(id string) (Account, error) {
	var name, email string

	err := a.db.QueryRow("SELECT name, email FROM accounts WHERE id = $1", id).
		Scan(&name, &email)

	if err != nil {
		return Account{}, err
	}

	return Account{ID: id, Name: name, Email: email}, nil
}
