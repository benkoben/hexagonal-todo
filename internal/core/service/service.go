package todo

import (
	"time"

	"github.com/benkoben/hexagonal-todo/internal/core/port"
)

type TodoServiceOptions struct {
    timeOut time.Duration
}

/**
  TodoService implements the ports for ListRepository, TaskRepository and SubTaskRepository.
  Its purpose is to provide access to lists, task and subtask repository systems.
*/
type TodoService struct {
   listRepo *port.ListRepository 
   taskRepo *port.TaskRepository 
   subtaskRepo *port.SubTaskRepository 
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
