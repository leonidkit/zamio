package service

import (
	"context"
	"zamio/internal/repository"
	"zamio/internal/service/payments"
)

type Payments interface {
	ProcessTransaction(ctx context.Context, firstEmail string, secondEmail string, sum int) error
}

type Services struct {
	Payments Payments
}

func New(repo repository.Interface) *Services {
	return &Services{
		Payments: payments.New(repo),
	}
}
