package postgres

import (
	"context"
	"log"
	"os"
	"testing"
	"zamio/internal/domain"
)

var (
	acc1 = domain.Account{
		ID:      123,
		Email:   "some@mail.ru",
		Balance: 100,
	}
	acc2 = domain.Account{
		ID:      123,
		Email:   "another@mail.ru",
		Balance: 2000,
	}
	ctx  = context.Background()
	repo = &DB{}
)

func TestMain(m *testing.M) {
	var err error

	repo, err = New("127.0.0.1", "5432", "admin", "admin", "zamio", false)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer repo.Close(ctx)

	_, err = repo.conn.Exec(ctx, "DELETE FROM accounts")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	_, err = repo.conn.Exec(ctx, "ALTER SEQUENCE accounts_id_seq RESTART WITH 1")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestCreateAccount(t *testing.T) {
	err := repo.CreateAccount(ctx, acc1)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	err = repo.CreateAccount(ctx, acc2)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}
}

func TestAccountByEmail(t *testing.T) {
	acc, err := repo.AccountByEmail(ctx, "some@mail.ru")
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	if acc.Email != acc1.Email {
		t.Fatalf("want user with email = %s, but recived email = %s", acc1.Email, acc.Email)
	}
}

func TestUpdateAccountsBalance(t *testing.T) {
	acc1.Balance -= 100
	acc2.Balance += 100

	err := repo.UpdateAccountsBalance(ctx, acc1, acc2)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	acc, err := repo.AccountByEmail(ctx, acc1.Email)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	if acc.Balance != acc1.Balance {
		t.Fatalf("want %s user with balance = %d, but recived balance = %d", acc1.Email, acc1.Balance, acc.Balance)
	}

	acc, err = repo.AccountByEmail(ctx, acc2.Email)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	if acc.Balance != acc2.Balance {
		t.Fatalf("want %s user with balance = %d, but recived balance = %d", acc2.Email, acc2.Balance, acc.Balance)
	}
}
