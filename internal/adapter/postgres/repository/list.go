package repository

import (
	"context"
	"fmt"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/adapter/postgres"
)

// TODO: Do I need to account for any context.Deadlines?

type ListRepository struct {
    db *postgres.DB
}

func NewListRepository(db *postgres.DB) *ListRepository {
    return &ListRepository{
        db,
    }
}

func (lr *ListRepository)CreateList(ctx context.Context, list *domain.List) (*domain.List, error) {

}
