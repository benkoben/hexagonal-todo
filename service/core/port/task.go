package port

import (
    "context"

    "github.com/benkoben/hexagonal-todo/service/core/domain"
)

// TaskRepository is an interface for interacting with our task related data
type TaskRepository interface {
    // Inserts a list item to the database
	CreateTask(ctx context.Context, list *domain.List) (*domain.List, error)

    // Retrieves a list item from the database by id
	GetTaskById(ctx context.Context, id int64) (*domain.List, error)

    // List retrieves all lists items from the database
    ListTasks(ctx context.Context) ([]*domain.List, error)

    // Removes a list item from the database by id
	DeleteTaskById(ctx context.Context, id int64)(*domain.List, error)
}


// TaskService is an interface for interacting with our task business logic
type TaskService interface {
    // Creates a new list 
	CreateTask(ctx context.Context, list *domain.List) (*domain.List, error)

    // Get a list by id
	GetTask(ctx context.Context, id int64) (*domain.List, error)

    // Retrieve all existing lists
    ListTasks(ctx context.Context) ([]*domain.List, error)

    // Delete a list by id
	DeleteTask(ctx context.Context, id int64)()
}
