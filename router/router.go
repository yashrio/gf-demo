package router

import (
    "gf-demo/app/api/hello"
	"gf-demo/app/service"
	"gf-demo/boot"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// jwt
	GfToken := &gtoken.GfToken{
		CacheMode:        g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		LoginPath:        "/auth/login",
		LoginBeforeFunc:  service.Auth.Login,
		LoginAfterFunc:   service.Auth.LoginAfter,
		LogoutPath:       "/auth/logout",
		AuthPaths:        g.SliceStr{"/system/*"},
		AuthAfterFunc:    service.Auth.AuthAfterFunc,
		LogoutBeforeFunc: service.Auth.LoginOut,
	}

	boot.GfToken = GfToken

	s.Group("/", func(group *ghttp.RouterGroup) {
		GfToken.Middleware(group)
		group.Middleware(service.MiddleWareService.Ctx)
		group.ALL("/hello", hello.Hello)
	})
}
