package main

import (
	"github.com/gin-gonic/gin"
	"heychart/router"
)

var port string
func init() {
	// init configure
	port = ":9090"

}
var quit = make(chan string)
func main() {
	route := gin.Default()
    go router.RegisterRoute(route)
	route.Run(port)
	<-quit
}
