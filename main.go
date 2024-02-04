package main

// main is the entrypoint for the application
// it initializes the configuration for the adapters and the todo service
// and starts the server

import (
	"log"
    "os"
    "context"

	"github.com/benkoben/hexagonal-todo/config"
	"github.com/benkoben/hexagonal-todo/internal/adapter/repository"
	"github.com/benkoben/hexagonal-todo/internal/adapter/handler/http"
	"github.com/benkoben/hexagonal-todo/internal/adapter/repository/postgres"
	todo "github.com/benkoben/hexagonal-todo/internal/core/service"
)

func main(){
    cfg, err := config.NewConfiguration()
    if err != nil {
        log.Fatalf("could not create configuration: %v", err)
    }

    // Initialize the database
    dbOptions := &postgres.PostgresClientOptions{
        Host:     cfg.Database.Host,
        Port:     cfg.Database.Port,
        Username: cfg.Database.User,
        Password: cfg.Database.Password,
        Database: cfg.Database.Database,
    }
    dbContext := context.Background()
    db, err := postgres.NewDB(dbContext, dbOptions)
    if err != nil {
        log.Fatalf("could not create database: %v", err)
    }

    // Initialize the List Repository
    repoOptions := repository.RepositoryOptions{
        Timeout: 0,
    }
    
    // Depenenency Injection
    // List
    listRepository := repository.NewListRepository(db, repoOptions)
    // Initialize the List service
    listService, err := todo.NewListService(listRepository) 
    if err != nil {
        log.Fatalf("could not create list service: %v", err)
    }
    listHandler := http.NewListHandler(listService)

    // Task
    // TODO: taskRepository := respository.NewTaskRepository(db, repoOptions)
    // TODO: Initialize the Task service
    // taskService, err := todo.NewListService(listRepository) 
    // if err != nil {
    //     log.Fatalf("could not create list service: %v", err)
    // }

    // Subtask
    // TODO: subtaskRepository := respository.NewSubtaskRepository(db, repoOptions)
    // TODO: Initialize the Subtask service
    // subtaskService, err := todo.NewListService(listRepository) 
    // if err != nil {
    //     log.Fatalf("could not create list service: %v", err)
    // }
    
    // Initialize the HTTP router
    router, err := http.NewRouter(*listHandler)
    if err != nil {
        log.Fatalf("could not create http router: %v", err)
    }

    // Intialize the server
    serverOpts := http.ServerOptions{
        Address: cfg.Server.Addr,
        Port: cfg.Server.Port,
        Router: router.Engine,
        Log: log.New(os.Stdout, "http: ", log.LstdFlags),
    }
    server, err := http.NewServer(serverOpts)
    if err != nil {
        log.Fatalf("could not create http server: %v", err)
    }

    // Start the server
    server.Start()

}
