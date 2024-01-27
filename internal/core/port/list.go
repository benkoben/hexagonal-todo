package port

import (
    "context"

    "github.com/benkoben/hexagonal-todo/internal/core/domain"
)


// ListRepository is an interface for interacting with our todo related data
type ListRepository interface {
    // Inserts a list item to the database
	CreateList(ctx context.Context, list *domain.List) (*domain.List, error)

    // Retrieves a list item from the database by id
	GetListById(ctx context.Context, id int64) (*domain.List, error)

    // List retrieves all lists items from the database
    GetLists(ctx context.Context) ([]*domain.List, error)

    // UpdateListById modifies an existing list in the database
    UpdateListById(ctx context.Context, id int64, updateAttrs domain.List)(*domain.List, error)

    // Removes a list item from the database by id
	DeleteListById(ctx context.Context, id int64)(*domain.List, error)
}


// ListService is an interface for interacting with our todo business logic
type ListService interface {
    // Creates a new list 
	CreateList(ctx context.Context, list *domain.List) (*domain.List, error)

    // Get a list by id
	GetListById(ctx context.Context, id int64) (*domain.List, error)

    // Retrieve all existing lists
    GetLists(ctx context.Context) ([]*domain.List, error)

    // Update an existing List 
    UpdateList(ctx context.Context, id int64, updateAttrs domain.List)(*domain.List, error)

    // Delete a list by id
	DeleteList(ctx context.Context, id int64)(*domain.List, error)
}
