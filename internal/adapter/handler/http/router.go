package http

import (
	service "github.com/benkoben/hexagonal-todo/internal/core/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
    *gin.Engine
}

func NewRouter(
        listService service.ListService,
    ) (*Router, error) {

    // TODO: learn CORS and implement in router
    router := gin.Default()
    v1 := router.Group("/v1") 
    v1.GET("/list/:id", listService.GetListById)
    v1.GET("/lists", listService.GetLists)
    v1.POST("/list", listService.CreateList)

    return &Router{
        router,
    }, nil
}
