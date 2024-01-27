package repository

import (
	"context"
	"time"

	"github.com/benkoben/hexagonal-todo/internal/adapter/repository/postgres"
	"github.com/benkoben/hexagonal-todo/internal/core/domain"
)

const (
    defaultTimeout = time.Second * 10
)

type ListRepositoryOptions struct {
    // Timeout sets the amount of time a query can take before context.WithTimeout times out.
    Timeout time.Duration
}

// TODO: Do I need to account for any context.Deadlines?
type ListRepository struct {
    db *postgres.DB
    queryTimeout time.Duration
}

func NewListRepository(db *postgres.DB, o ListRepositoryOptions) *ListRepository {
    if o.Timeout == 0 {
        o.Timeout = defaultTimeout
    }

    return &ListRepository{
        db: db,
        queryTimeout: o.Timeout,
    }
}

func (lr *ListRepository)CreateList(ctx context.Context, list *domain.List) (*domain.List, error) {
    // Build QueryStatement 
    ctx, cancel := context.WithTimeout(ctx, lr.queryTimeout)
    defer cancel()

    listStatement := lr.db.QueryBuilder.Insert("List").
                        Columns("name", "created_at").
                        Values(list.Name, list.CreatedAt).
                        Suffix("RETURNING *")

    // Convert to SQL string
    query, _, err := listStatement.ToSql()
    if err != nil {
        return nil, err
    }

    // Run Query on database
    lr.db.QueryRowEx(ctx, query, nil, nil).Scan(list)
    // Scan and update list

    // Return list
    return list, nil
}
