package main

import (
	"github.com/gin-gonic/gin"
	"heychart/router"
)

var quit = make(chan string)
func main() {
	route := gin.Default()
    go router.RegisterRoute(route)
	route.Run(":9090")
	<-quit
}
