package router

import "github.com/gogf/gf/v2/net/ghttp"

// MiddlewareCORS 解决跨域问题
func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"*"}
	//corsOptions.AllowOrigin="http://localhost:8900"
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
