package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

var hub *Hub

func init() {
	hub = newHub()
	go hub.run()
}

type Result struct {
	Code    int
	State   bool
	Message string
	Data    interface{}
}

func (r *Result) SetCode(code int) {
	r.Code = code
}

func (r *Result) SetResult(code int, status bool, message string, data interface{}) {
	r.Code = code
	r.State = status
	r.Message = message
	r.Data = data
}

func TestFunction(request *gin.Context) {

	var resp Result
	resp.SetResult(200, true, "Hello! get data successful", "gogogogo")
	request.JSON(200, resp)

}

func WebsocketServer(r *gin.Context) {
	serveWs(hub, r.Writer, r.Request)

}

func Home(r *gin.Context) {
	log.Println(r.Request.URL)
	if r.Request.URL.Path != "/" {
		http.Error(r.Writer, "Not found", http.StatusNotFound)
		return
	}
	if r.Request.Method != "GET" {
		http.Error(r.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(r.Writer, r.Request, "home.html")
}
