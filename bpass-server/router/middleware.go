package router

import "github.com/gogf/gf/net/ghttp"

// MiddlewareCORS 解决跨域问题
func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"*"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
