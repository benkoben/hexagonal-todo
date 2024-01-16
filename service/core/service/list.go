package todo

import (
	"context"
	"time"

	"github.com/benkoben/hexagonal-todo/core/domain"
	"github.com/benkoben/hexagonal-todo/core/port"
	"github.com/benkoben/hexagonal-todo/core/util"
)

const (
    defaultServiceTimeOut = time.Second * 30
)

type TodoServiceOptions struct {
    timeOut time.Duration
}

/**
  TodoService implements the ports for ListRepository, TaskRepository and SubTaskRepository.
  Its purpose is to provide access to lists, task and subtask repository systems.
*/
type TodoService struct {
   listRepo port.ListRepository 
   taskRepo port.TaskRepository 
   subtaskRepo port.SubTaskRepository 
   timeout time.Duration
}

/*
    NewTodoService() is a contructor for the TodoService
*/
func NewTodoService(opts *TodoServiceOptions, listRepo *port.ListRepository, taskRepo *port.TaskRepository, subTaskRepo *port.SubTaskRepository) *TodoService {

    if opts.timeOut == 0 {
       opts.timeOut = defaultServiceTimeOut 
    }

    return &TodoService{
        listRepo: listRepo,
        taskRepo: taskRepo,
        subtaskRepo: subTaskRepo,
    }
}

func (t *TodoService) CreateList(ctx context.Context, list *domain.List) (*domain.List, error){
    createdAt := time.Now()
    list.createdAt = createdAt
    
    deadline := time.Now().Add(time.Second * t.timeout)
    ctx, cancelCtx := context.WithDeadline(ctx, deadline)
    defer cancelCtx()
    
    _, err := t.CreateList(ctx,list)
    if err != nil {
        return list, err
    }

    return list, nil
}

func (t *TodoService) GetListById(ctx context.Context, id int64) (*domain.List, error){
    deadline := time.Now().Add(time.Second * t.timeout)
    ctx, cancelCtx := context.WithDeadline(ctx, deadline)
    defer cancelCtx()

    list, err := t.GetListById(ctx, id) 
    if err != nil {
        return domain.List{}, err
    }

    return list, nil
}


func (t *TodoService) ListLists(ctx context.Context) ([]*domain.List, error){

    deadline := time.Now().Add(time.Second * t.timeout)
    ctx, cancelCtx := context.WithDeadline(ctx, deadline)
    defer cancelCtx()

    lists, err := t.ListLists(ctx)
    if err != nil {
        return []domain.List{}, err
    }

    return lists, nil
}


func (t *TodoService) DeleteListById(ctx context.Context, id int64)(*domain.List, error){
    
    deadline := time.Now().Add(time.Second * t.timeout)
    ctx, cancelCtx := context.WithDeadline(ctx, deadline)
    defer cancelCtx()

    list, err := t.DeleteListById(ctx, id) 
    if err != nil {
        return domain.List{}, err
    }

    return list, nil
}
