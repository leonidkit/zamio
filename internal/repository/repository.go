package repository

import (
	"context"
	"zamio/internal/domain"
)

// AccountRepository is interface defining the contract for working with the account
type AccountRepository interface {
	// Create add account with empty balance
	CreateAccount(ctx context.Context, account domain.Account) error

	// ByEmail return account, takes user's email
	AccountByEmail(ctx context.Context, email string) (domain.Account, error)

	// UpdateAccountsBalance updates two account's balance via a transaction
	UpdateAccountsBalance(ctx context.Context, accountFirst, accountSecond domain.Account) error
}

// Interface is interface defining the contract for working with the app database
type Interface interface {
	AccountRepository
}
