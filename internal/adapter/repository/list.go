package repository

import (
	"context"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/adapter/repository/postgres"
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

    // Build QueryStatement 
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
    lr.db.QueryRow(query).Scan(list)
    // Scan and update list

    // Return list
    return list, nil
}
