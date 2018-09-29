package router

import (
	"github.com/gin-gonic/gin"
	"vesyncmqtt/src/github.com/robvdl/pongo2gin"
	"heychart/controller"
)

func RegisterRoute(route *gin.Engine) {
	opt := pongo2gin.RenderOptions{
		TemplateDir: "templates/",
		//TemplateDir: "E:/linus/code/gitlab/vesyncmqtt/src/templates/",
	}

	route.HTMLRender = pongo2gin.New(opt)
	route.GET("/",controller.TestFunction)


}
