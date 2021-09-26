package transaction

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for transactions.
type Service interface {
	Get(ctx context.Context, id string) (Transaction, error)
	Query(ctx context.Context, offset, limit int) ([]Transaction, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input entity.TransactionRequest) (Transaction, error)
}

// Transaction represents the data about an transaction.
type Transaction struct {
	entity.TransactionResponse
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new transaction service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the transaction with the specified the transaction ID.
func (s service) Get(ctx context.Context, id string) (Transaction, error) {
	transaction, err := s.repo.Get(ctx, id)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{transaction}, nil
}

// Create creates a new transaction.
func (s service) Create(ctx context.Context, req entity.TransactionRequest) (Transaction, error) {
	if err := req.Validate(); err != nil {
		return Transaction{}, err
	}
	id := entity.GenerateID()

	err := s.repo.Create(ctx, entity.TransactionRequest{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
	})
	if err != nil {
		return Transaction{}, err
	}
	return s.Get(ctx, id)
}

// Count returns the number of transactions.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the transactions with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Transaction, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Transaction{}
	for _, item := range items {
		result = append(result, Transaction{item})
	}
	return result, nil
}
