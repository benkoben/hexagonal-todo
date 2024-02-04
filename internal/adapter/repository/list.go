package repository

import (
	"context"
	"time"
    "fmt"

	"github.com/benkoben/hexagonal-todo/internal/adapter/repository/postgres"
	"github.com/benkoben/hexagonal-todo/internal/core/domain"
)

const (
    defaultTimeout = time.Second * 10
)

type ListRepository struct {
    db *postgres.DB
    queryTimeout time.Duration
}

func NewListRepository(db *postgres.DB, o RepositoryOptions) *ListRepository {
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

    listStatement := lr.db.QueryBuilder.Insert("list").
                        Columns("name", "created_at").
                        Values(list.Name, list.CreatedAt).
                        Suffix("RETURNING *")

    // Convert to SQL string
    query, _, err := listStatement.ToSql()
    if err != nil {
        return nil, err
    }

    // Run Query on database
    fmt.Printf("Query: %s\n", query)
    rows, err := lr.db.Pool.Query(ctx, query, list.Name, list.CreatedAt)
    if err != nil {
        return nil, fmt.Errorf("could not create list: %v", err)
    }
    defer rows.Close()
    rows.Scan(&list.Id, &list.Name, &list.CreatedAt)

    // Return list
    return list, nil
}

// TODO implement following methods
func (lr *ListRepository)GetListById(ctx context.Context, id int64) (*domain.List, error) {return nil, nil}

func (lr *ListRepository)GetLists(ctx context.Context) ([]*domain.List, error) {return nil, nil}

func (lr *ListRepository)UpdateListById(ctx context.Context, id int64, updateAttrs domain.List) (*domain.List, error) {return nil, nil}

func (lr *ListRepository)DeleteListById(ctx context.Context, id int64) (*domain.List, error) {return nil, nil}
