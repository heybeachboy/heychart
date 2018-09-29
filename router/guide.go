package router

import (
	"github.com/gin-gonic/gin"
	"heychart/controller"
)

func RegisterRoute(route *gin.Engine) {
	 route.GET("/",controller.TestFunction)


}
