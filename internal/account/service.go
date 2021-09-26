package account

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for accounts.
type Service interface {
	Get(ctx context.Context, id string) (Account, error)
	Query(ctx context.Context, offset, limit int) ([]Account, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input entity.AccountRequest) (Account, error)
}

// Account represents the data about an account.
type Account struct {
	entity.AccountResponse
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new account service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the account with the specified the account ID.
func (s service) Get(ctx context.Context, id string) (Account, error) {
	account, err := s.repo.Get(ctx, id)
	if err != nil {
		return Account{}, err
	}
	return Account{account}, nil
}

// Create creates a new account.
func (s service) Create(ctx context.Context, req entity.AccountRequest) (Account, error) {
	if err := req.Validate(); err != nil {
		return Account{}, err
	}
	id := entity.GenerateID()

	err := s.repo.Create(ctx, entity.AccountRequest{
		ID:             id,
		DocumentNumber: req.DocumentNumber,
	})
	if err != nil {
		return Account{}, err
	}
	return s.Get(ctx, id)
}

// Count returns the number of accounts.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the accounts with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Account, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Account{}
	for _, item := range items {
		result = append(result, Account{item})
	}
	return result, nil
}
