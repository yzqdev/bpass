package router

import (
	"b0pass/apps/api"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gookit/color"
)

func init() {
	s := g.Server()

	// Chat
	s.BindHandler("/ws", func(r *ghttp.Request) {
		var ctx = r.Context()
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(ctx, err)
			r.Exit()
		}

		for {
			msgType, msg, err := ws.ReadMessage()
			fmt.Println(string(msg))
			if err != nil {
				return
			}

			if err = ws.WriteMessage(msgType, msg); err != nil {
				color.Redln("返回值错误")
				return
			}
		}
	})
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.RedirectTo("/spa/")
	})
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
