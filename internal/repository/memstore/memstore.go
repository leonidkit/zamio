package memstore

import (
	"context"
	"fmt"
	"sync"
	"zamio/internal/domain"
)

type DB struct {
	mux  *sync.RWMutex
	accs []domain.Account
}

func New() *DB {
	return &DB{
		mux:  &sync.RWMutex{},
		accs: make([]domain.Account, 0, 10),
	}
}

func (db *DB) CreateAccount(ctx context.Context, acc domain.Account) (err error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	db.accs = append(db.accs, acc)
	return
}

func (db *DB) UpdateAccountsBalance(ctx context.Context, accountFirst, accountSecond domain.Account) error {
	db.mux.RLock()
	defer db.mux.RUnlock()

	for i, u := range db.accs {
		if u.Email == accountFirst.Email {
			db.accs[i] = accountFirst
		}
		if u.Email == accountSecond.Email {
			db.accs[i] = accountSecond
		}
	}
	return nil
}

func (db *DB) AccountByEmail(ctx context.Context, email string) (domain.Account, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	for i, u := range db.accs {
		if u.Email == email {
			db.accs[i].Email = email
			return u, nil
		}
	}

	return domain.Account{}, fmt.Errorf("user with email %s not found", email)
}
