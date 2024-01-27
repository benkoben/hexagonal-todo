package http


func (s server) routes() {
    s.router.GET("/lists", getLists)
    s.router.POST("/lists", createList)
}
