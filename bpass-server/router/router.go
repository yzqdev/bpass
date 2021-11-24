package router

import (
	"b0pass/apps/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gookit/color"
)

func init() {
	s := g.Server()

	// Index
color.Redln("刀刀赚了")
	// Chat
	//s.BindHandler("/chat", new(chat.Controller))
	//s.BindHandler("/sync ", api.SyncIndex)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/fileList", api.FileLists)
			group.GET("/index", api.Index)
		})

		group.Group("/api", func(g *ghttp.RouterGroup) {
			//cors

			//file
			g.POST("/upload", api.Upload)
			g.GET("/lists", api.Lists)
			g.GET("/delete", api.Delete)
			g.GET("/dump", api.Dump)
			g.GET("/upload", api.UploadShow)
			//server
			g.GET("/globalData", api.GlobalData)
			g.ALL("/subpath", api.GetSubPath)
			g.ALL("/textdata", api.GetTextData)
			g.GET("/openurl", api.OpenUrl)

		})
	})
	// Api

}
