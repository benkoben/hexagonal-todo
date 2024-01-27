package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
    service "github.com/benkoben/hexagonal-todo/internal/core/service"
)

const (
    defaultAddr = "0.0.0.0"
    defaultPort = "8080"
)

type logger interface {
    Printf(format string, v...any)
    Println(v...any)
    Fatalf(format string, v...any)
    Fatalln(v...any)
}

type ServerOptions struct {
    Address string
    Port string
    Router *gin.Engine
    Log logger
}

type server struct {
    // Which interface/address the service should listen on
    address string
    // Port that the server should listen to
    port string
    // Routes which the server should handle
    router *gin.Engine
    // Logger implements function for writing logs 
    log logger
    // Todo implements the core service for our todo app
    todo *service.TodoService
}

func NewServer(o ServerOptions) (*server, error) {
    if o.Address == "" {
        o.Address = defaultAddr
    }

    if o.Port == "" {
        o.Port = defaultPort
    }

    if o.Router == nil {
        return nil, fmt.Errorf("router field cannot be nil in http server options")
    }

    if o.Log == nil {
        return nil, fmt.Errorf("log field cannot be nil in http server options")
    }

    return &server{
        address: o.Address,
        port: o.Port,
        router: o.Router,
    }, nil
}

func (s server)startServer() error {

}

func (s server)stopServer() error {

}
