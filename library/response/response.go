package response

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

const (
	SuccesCode 	= 0
	ErrorCode 	= -1
)

type JsonResponse struct {
	Code 	int			`json:"code"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

var response = new(JsonResponse)

func Json(r *ghttp.Request, code int, message string, data ...interface{})  {
	fmt.Println(data)
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code: 		code,
		Message: 	message,
		Data: 		responseData,
	})
}

func JsonExit(r *ghttp.Request, err int, message string, data ...interface{})  {
	Json(r, err, message, data)
	r.Exit()
}

func SuccExit(r *ghttp.Request, message string, data ...interface{})  {
	Json(r, SuccesCode, message, data)
	r.Exit()
}

func Succ(r *ghttp.Request, message string, data ...interface{})  {
	Json(r, SuccesCode, message, data)
}

func FailExit(r *ghttp.Request, err int, message string, data ...interface{})  {
	Json(r, err, message, data)
}

func Fail(r *ghttp.Request, err int, message string, data ...interface{})  {
	Json(r, err, message, data)
	r.Exit()
}