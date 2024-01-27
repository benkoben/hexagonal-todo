package todo

import (
	"context"
	"time"
    "fmt"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/core/port"
)

type SubtaskServiceOptions struct {
    // Timeout is used to control how long a backend query can take before it times out 
    Timeout time.Duration
}

type SubtaskService struct {
    // Interface which implements accessing data in list repository
    SubtaskRepo port.SubTaskRepository
    // Timeout controls how long a query to the repository should take at most
    Timeout time.Duration
}

func (s *SubtaskService) CreateSubtask(ctx context.Context, list *domain.Subtask) (*domain.Subtask, error){
    createdAt := time.Now()
    list.CreatedAt = createdAt
   
    _, err := s.SubtaskRepo.CreateSubtask(ctx,list)
    if err != nil {
        return list, fmt.Errorf("could not create list: %v", err)
    }

    return list, nil
}

func (s *SubtaskService) GetSubtaskById(ctx context.Context, id int64) (*domain.Subtask, error){
    list, err := s.SubtaskRepo.GetSubtaskById(ctx, id) 
    if err != nil {
        return &domain.Subtask{}, err
    }

    return list, nil
}


func (s *SubtaskService) GetSubtasks(ctx context.Context) ([]*domain.Subtask, error){
    lists, err := s.SubtaskRepo.GetSubtasks(ctx)
    if err != nil {
        return []*domain.Subtask{}, err
    }

    return lists, nil
}

func (s *SubtaskService) DeleteSubtaskById(ctx context.Context, id int64)(*domain.Subtask, error){

    list, err := s.DeleteSubtaskById(ctx, id) 
    if err != nil {
        return &domain.Subtask{}, err
    }

    return list, nil
}
