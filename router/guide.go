package router

import (
	"github.com/gin-gonic/gin"
	"heychart/controller"
	"github.com/gin-contrib/cors"
)

func RegisterRoute(route *gin.Engine) {
	 route.LoadHTMLFiles("template/**/*")
	 cfg := cors.DefaultConfig()
	 cfg.AllowAllOrigins = true
	 route.Use(cors.Default())
	 route.GET("/",controller.Home)
	 route.GET("/ws",controller.WebSocketServer)
	// route.GET("/",controller.TestFunction)


}
