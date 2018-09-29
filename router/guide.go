package router

import (
	"github.com/gin-gonic/gin"
	"heychart/controller"
)

func RegisterRoute(route *gin.Engine) {
	 route.LoadHTMLFiles("template/**/*")
	 route.GET("/",controller.Home)
	 route.GET("/ws",controller.WebsocketServer)
	// route.GET("/",controller.TestFunction)


}
