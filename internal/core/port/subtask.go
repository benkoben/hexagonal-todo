package port

import (
    "context"

    "github.com/benkoben/hexagonal-todo/internal/core/domain"
)

// SubTaskRepository is an interface for interacting with our subtask related data
type SubTaskRepository interface {
    // Inserts a list item to the database
	CreateTask(ctx context.Context, list *domain.List) (*domain.List, error)

    // Retrieves a list item from the database by id
	GetTaskById(ctx context.Context, id int64) (*domain.List, error)

    // List retrieves all lists items from the database
    GetTasks(ctx context.Context) ([]*domain.List, error)

    // UpdateListById modifies an existing list in the database
    UpdateSubTaskById(ctx context.Context, id int64, updateAttrs domain.Subtask)(*domain.Subtask, error)

    // Removes a list item from the database by id
	DeleteTaskById(ctx context.Context, id int64)(*domain.List, error)
}


// SubTaskService is an interface for interacting with our subtask business logic
type SubTaskService interface {
    // Creates a new list 
	CreateTask(ctx context.Context, list *domain.List) (*domain.List, error)

    // Get a list by id
	GetTask(ctx context.Context, id int64) (*domain.List, error)

    // Retrieve all existing lists
    GetTasks(ctx context.Context) ([]*domain.List, error)

    // Update an existing List 
    UpdateSubTask(ctx context.Context, id int64, updateAttrs domain.Subtask)(*domain.Subtask, error)

    // Delete a list by id
	DeleteTask(ctx context.Context, id int64)(*domain.Task, error)
}
