package port

import (
    "context"

    "github.com/benkoben/hexagonal-todo/internal/core/domain"
)

// SubTaskRepository is an interface for interacting with our subtask related data
type SubTaskRepository interface {
    // Inserts a list item to the database
	CreateSubtask(ctx context.Context, list *domain.Subtask) (*domain.Subtask, error)

    // Retrieves a list item from the database by id
	GetSubtaskById(ctx context.Context, id int64) (*domain.Subtask, error)

    // List retrieves all lists items from the database
    GetSubtasks(ctx context.Context) ([]*domain.Subtask, error)

    // UpdateListById modifies an existing list in the database
    UpdateSubtaskById(ctx context.Context, id int64, updateAttrs domain.Subtask)(*domain.Subtask, error)

    // Removes a list item from the database by id
	DeleteSubTaskById(ctx context.Context, id int64)(*domain.Subtask, error)
}


// SubTaskService is an interface for interacting with our subtask business logic
type SubTaskService interface {
    // Creates a new list 
	CreateSubtask(ctx context.Context, list *domain.Subtask) (*domain.Subtask, error)

    // Get a list by id
	GetSubtask(ctx context.Context, id int64) (*domain.Subtask, error)

    // Retrieve all existing lists
    GetSubtasks(ctx context.Context) ([]*domain.Subtask, error)

    // Update an existing List 
    UpdateSubTask(ctx context.Context, id int64, updateAttrs domain.Subtask)(*domain.Subtask, error)

    // Delete a list by id
	DeleteSubtask(ctx context.Context, id int64)(*domain.Task, error)
}
