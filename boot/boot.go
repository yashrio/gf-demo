package boot

import (
	"gf-demo/library/wechat"
	_ "gf-demo/packed"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/swagger"
)

var GfToken *gtoken.GfToken

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	g.SetDebug(true)

	// 使用Redis缓存中间件
	cache := gcache.New()
	adapter := adapter.NewRedis(g.Redis())
	cache.SetAdapter(adapter)

	// 配置微信
	wechat.InitWechat()
}
