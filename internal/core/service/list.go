package todo

import (
	"context"
	"time"
    "fmt"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/core/port"
)


type ListService struct {
    // Interface which implements accessing data in list repository
    ListRepo port.ListRepository
    // Timeout controls how long a query to the repository should take at most
    Timeout time.Duration
}

func NewListService(listRepo port.ListRepository) (*ListService, error){
    return &ListService{
        ListRepo: listRepo,
    }, nil
}

func (s *ListService) CreateList(ctx context.Context, list *domain.List) (*domain.List, error){
    createdAt := time.Now()
    list.CreatedAt = createdAt
   
    _, err := s.ListRepo.CreateList(ctx,list)
    if err != nil {
        return list, fmt.Errorf("could not create list: %v", err)
    }

    return list, nil
}

func (s *ListService) GetListById(ctx context.Context, id int64) (*domain.List, error){
    list, err := s.ListRepo.GetListById(ctx, id) 
    if err != nil {
        return &domain.List{}, fmt.Errorf("GetListById could not fetch list from repository: %v", err)
    }

    return list, nil
}


func (s *ListService) GetLists(ctx context.Context) ([]*domain.List, error){
    lists, err := s.ListRepo.GetLists(ctx)
    if err != nil {
        return []*domain.List{}, err
    }

    return lists, nil
}

func (s *ListService) UpdateList(ctx context.Context, id int64, updateAttrs domain.List)(*domain.List, error){

    list, err := s.ListRepo.UpdateListById(ctx, id, updateAttrs) 
    if err != nil {
        return &domain.List{}, err
    }

    return list, nil
}

func (s *ListService) DeleteList(ctx context.Context, id int64)(*domain.List, error){

    list, err := s.ListRepo.DeleteListById(ctx, id) 
    if err != nil {
        return &domain.List{}, err
    }

    return list, nil
}
