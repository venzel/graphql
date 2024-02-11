package graph

import "graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AccountDB     *database.Account
	TransactionDB *database.Transaction
}
