package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
    *gin.Engine
}

func NewRouter(
        listHandler ListHandler,
    ) (*Router, error) {

    // TODO: learn CORS and implement in router
    router := gin.Default()
    v1 := router.Group("/v1") 
    v1.GET("/list/:id", listHandler.GetListById)
    v1.GET("/lists", listHandler.GetLists)
    v1.POST("/list", listHandler.CreateList)

    return &Router{
        router,
    }, nil
}
