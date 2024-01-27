package port

import (
    "context"

    "github.com/benkoben/hexagonal-todo/internal/core/domain"
)

// TaskRepository is an interface for interacting with our task related data
type TaskRepository interface {
    // Inserts a list item to the database
	CreateTask(ctx context.Context, list *domain.Task) (*domain.Task, error)

    // Retrieves a list item from the database by id
	GetTaskById(ctx context.Context, id int64) (*domain.Task, error)

    // List retrieves all lists items from the database
    GetTasks(ctx context.Context) ([]*domain.Task, error)

    // UpdateListById modifies an existing list in the database
    UpdateTaskById(ctx context.Context, id int64, updateAttrs domain.Task)(*domain.Task, error)

    // Removes a list item from the database by id
	DeleteTaskById(ctx context.Context, id int64)(*domain.Task, error)
}


// TaskService is an interface for interacting with our task business logic
type TaskService interface {
    // Creates a new list 
	CreateTask(ctx context.Context, list *domain.Task) (*domain.Task, error)

    // Get a list by id
	GetTask(ctx context.Context, id int64) (*domain.Task, error)

    // Retrieve all existing lists
    ListTasks(ctx context.Context) ([]*domain.Task, error)

    // Update an existing List 
    UpdateTask(ctx context.Context, id int64, updateAttrs domain.Task)(*domain.Task, error)

    // Delete a list by id
	DeleteTask(ctx context.Context, id int64)()
}
