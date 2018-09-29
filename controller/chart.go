package controller

import "github.com/gin-gonic/gin"

type Result struct {
	 Code int
	 State bool
	 Message string
	 Data interface{}
}

func(r *Result)SetCode(code int) {
	r.Code = code
}

func(r *Result)SetResult(code int, status bool, message string, data interface{}) {
	r.Code = code
	r.State = status
	r.Message = message
	r.Data = data
}
func TestFunction(request *gin.Context) {

	var resp Result
	 resp.SetResult(200,true,"Hello! get data successful","gogogogo")
	 request.JSON(200,resp)

}

