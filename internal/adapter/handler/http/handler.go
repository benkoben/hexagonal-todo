package http

// TODO: Add logger so that we can catch issues on the backend

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// response wraps around JSON and Code.
type response interface {
    JSON() []byte
    Code() int
}

// write sends a response to the client
func write(w http.ResponseWriter, response response) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(response.Code())
	w.Write(response.JSON())
}

func (s server) createList(c *gin.Context){
    var list domain.List
    ctx := context.Background()
    // Marshal payload into a domain.List struct
    if err := c.BindJSON(&list); err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    } 

    repositoryResponse, err := s.todo.CreateList(ctx, &list)
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    write(c.Writer, newResponse(*repositoryResponse, http.StatusOK))
} 

func getLists(c *gin.Context){
    return c.IndentedJSON(http.StatusOK, domain.List)
} 
