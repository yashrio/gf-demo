package main

import (
	_ "gf-demo/boot"
	_ "gf-demo/router"

	"github.com/gogf/gf/frame/g"
)

// @title       `gf-demo`示例服务API
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	g.Server().Run()
}
