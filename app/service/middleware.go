package service

import (
	"fmt"
	"gf-demo/app/model"
	"gf-demo/boot"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var MiddleWareService = new(middleWareService)

type middleWareService struct {}

// 初始化上下文管理器
func (m *middleWareService) Ctx(r *ghttp.Request)  {
	fmt.Println("ctx初始化")
	customCtx := &model.Context{
		User: nil,
	}
	Context.Init(r, customCtx)

	// 解析当前用户
	resp := boot.GfToken.GetTokenData(r)
	userInfo := new(model.ContextUser)
	gconv.Struct(resp.Get("data"), &userInfo)
	if userInfo != nil {
		customCtx.User = userInfo
	}
	r.Middleware.Next()
}
