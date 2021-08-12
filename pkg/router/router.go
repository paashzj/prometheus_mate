package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"prometheus_mate/pkg/api"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/v1/prometheus/", func(group *ghttp.RouterGroup) {
		group.Group("/jobs", func(group *ghttp.RouterGroup) {
			group.POST("/", api.AddJob)
			group.DELETE("/{job}", api.DelJob)
		})
	})
}
