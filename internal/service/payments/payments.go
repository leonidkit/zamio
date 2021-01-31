package payments

import (
	"context"
	"fmt"
	"zamio/internal/repository"
)

type PaymentService struct {
	repo repository.AccountRepository
}

func New(repo repository.AccountRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

// ProcessTransaction transfers sum money from the firstAcc account to the secondAcc one
func (p *PaymentService) ProcessTransaction(ctx context.Context, firstEmail string, secondEmail string, sum int) error {
	if firstEmail == secondEmail {
		return fmt.Errorf("it is not possible to transfer funds to yourself")
	}

	firstAcc, err := p.repo.AccountByEmail(ctx, firstEmail)
	if err != nil {
		return fmt.Errorf("account getting error: %v", err)
	}

	secondAcc, err := p.repo.AccountByEmail(ctx, secondEmail)
	if err != nil {
		return fmt.Errorf("account getting error: %v", err)
	}

	firstAcc.Balance -= sum
	secondAcc.Balance += sum

	err = p.repo.UpdateAccountsBalance(ctx, firstAcc, secondAcc)
	if err != nil {
		return fmt.Errorf("payment error: %v", err)
	}

	return err
}
