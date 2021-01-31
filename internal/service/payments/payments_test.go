package payments

import (
	"context"
	"testing"
	"zamio/internal/domain"
	"zamio/internal/repository/memstore"
)

var (
	repo = memstore.New()
	acc1 = domain.Account{
		ID:      123,
		Email:   "some@mail.ru",
		Balance: 100,
	}
	acc2 = domain.Account{
		ID:      124,
		Email:   "another@mail.ru",
		Balance: 100,
	}
	ctx = context.Background()
)

func TestPaymentService(t *testing.T) {
	sum := 100
	paymentSvc := New(repo)

	paymentSvc.repo.CreateAccount(ctx, acc1)
	paymentSvc.repo.CreateAccount(ctx, acc2)

	err := paymentSvc.ProcessTransaction(ctx, acc1.Email, acc2.Email, sum)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	acc, err := paymentSvc.repo.AccountByEmail(ctx, acc2.Email)
	if err != nil {
		t.Fatalf("get error, but not expected: %v", err)
	}

	if acc.Balance != (acc2.Balance + sum) {
		t.Fatalf("want %s user balance = %d, but recived balance = %d", acc2.Email, acc2.Balance+sum, acc.Balance)
	}
}
