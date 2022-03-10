package router

import (
	"b0pass/apps/api"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()

	// Chat
	//s.BindHandler("/chat", new(chat.SyncController))
	s.BindObject("/sync ", new(api.SyncController))
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/fileList", api.FileLists)
			group.GET("/index", api.Index)
			group.GET("/ips", api.Ips)
			group.POST("/upload", api.MultiUpload)
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
			g.POST("/textdata", api.SendTextData)
			g.GET("/textdata", api.GetTextData)
			g.GET("/openurl", api.OpenUrl)

		})
	})
	// Api

}
