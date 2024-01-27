package http

// TODO: Add logger so that we can catch issues on the backend

import (
	"context"
	"net/http"
	"strconv"

	"github.com/benkoben/hexagonal-todo/internal/core/domain"
	"github.com/benkoben/hexagonal-todo/internal/core/port"
	"github.com/gin-gonic/gin"
)

type ListHandler struct {
    svc port.ListService 
}

func (lh ListHandler)CreateList(c *gin.Context){
    var list domain.List
    ctx := context.Background()
    // Marshal payload into a domain.List struct
    if err := c.BindJSON(&list); err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    } 

    repositoryResponse, err := lh.svc.CreateList(ctx, &list)
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    write(c.Writer, newResponse(*repositoryResponse, http.StatusOK))
} 

/*
    getList checks if id parameter is present in the requested url before fetching a specific list from the listService.
    If there is not id parameter present then all existing lists will be fetched from the listService.
*/
func (lh ListHandler) GetListById(c *gin.Context){

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        // TODO:  Add logging entry here
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    ctx := context.Background()
    list, err := lh.svc.GetListById(ctx, id) 
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }
    write(c.Writer, newResponse(list, http.StatusOK))
}

func (lh ListHandler) GetLists(c *gin.Context) {
    ctx := context.Background()
    lists, err := lh.svc.GetLists(ctx) 
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }
    write(c.Writer, newResponse(lists, http.StatusOK))
}

func (lh ListHandler) UpdateList(c *gin.Context){
    var list domain.List
    ctx := context.Background()

    // Marshal payload into a domain.List struct
    if err := c.BindJSON(&list); err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        // TODO:  Add logging entry here
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    updatedList, err := lh.svc.UpdateList(ctx, id, list) 
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    write(c.Writer, newResponse(updatedList, http.StatusOK))
}

func (lh ListHandler) DeleteList(c *gin.Context){

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        // TODO:  Add logging entry here
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }

    ctx := context.Background()
    list, err := lh.svc.DeleteList(ctx, id) 
    if err != nil {
        write(c.Writer, newError(errInternalServer, http.StatusInternalServerError))
    }
    write(c.Writer, newResponse(list, http.StatusOK))
}
