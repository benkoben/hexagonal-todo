package todo

import (
	"context"
	"time"
    "fmt"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/core/port"
)

type TaskServiceOptions struct {
    // Timeout is used to control how long a backend query can take before it times out 
    Timeout time.Duration
}

type TaskService struct {
    // Interface which implements accessing data in list repository
    TaskRepo port.TaskRepository
    // Timeout controls how long a query to the repository should take at most
    Timeout time.Duration
}

func (s *TaskService) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error){
    createdAt := time.Now()
    task.CreatedAt = createdAt
   
    _, err := s.TaskRepo.CreateTask(ctx, task)
    if err != nil {
        return task, fmt.Errorf("could not create list: %v", err)
    }

    return task, nil
}

func (s *TaskService) GetTaskById(ctx context.Context, id int64) (*domain.Task, error){
    list, err := s.TaskRepo.GetTaskById(ctx, id) 
    if err != nil {
        return &domain.Task{}, err
    }

    return list, nil
}


func (s *TaskService) GetTasks(ctx context.Context) ([]*domain.Task, error){
    lists, err := s.TaskRepo.GetTasks(ctx)
    if err != nil {
        return []*domain.Task{}, err
    }

    return lists, nil
}

func (s *TaskService) DeleteTaskById(ctx context.Context, id int64)(*domain.Task, error){

    list, err := s.DeleteTaskById(ctx, id) 
    if err != nil {
        return &domain.Task{}, err
    }

    return list, nil
}


