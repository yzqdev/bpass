package boot

import (
	"b0pass/library/fileinfos"
	"flag"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

var (
	PathRoot string
	ServPort int
)
var ctx = gctx.New()

func ExecArgs() {
	flag.Parse()
	if ServPort <= 0 {
		env, _ := g.Cfg().GetWithEnv(ctx, "setting.port")
		ServPort = env.Int()
	}
}

// 用于应用初始化。
func init() {

	// 分析CLI参数
	flag.IntVar(&ServPort, "p", ServPort, "-p for Server Port(default=8899)")
	ExecArgs()

	// 资源根目录
	PathRoot = fileinfos.GetRootPath()

	// 恢复文件到缓存
	fileinfos.Init("data_path", "data_text")

	go func() {

		// APP核心引擎
		//template文件夹用于非spa
		//v := g.View()
		c := g.Config()
		s := g.Server()

		// 加载动作缓冲
		time.Sleep(3000 * time.Millisecond)

		// 模板引擎配置
		//_ = v.AddPath("template")
		//v.SetDelimiters("${", "}")

		// glog配置
		logpath, _ := c.Get(ctx, "setting.logpath")
		_ = glog.SetPath(logpath.String())
		glog.SetStdoutPrint(true)

		// Web Server配置
		s.SetIndexFolder(true)
		//s.SetAddr("0.0.0.0" + strconv.Itoa(ServPort))
		//s.SetServerRoot("public")
		s.SetLogPath(logpath.String())
		s.SetReadTimeout(3 * 60 * time.Second)
		s.SetWriteTimeout(3 * 60 * time.Second)
		s.SetIdleTimeout(3 * 60 * time.Second)
		s.SetMaxHeaderBytes(32 * 1024)
		s.SetErrorLogEnabled(true)
		s.SetAccessLogEnabled(true)
		s.SetPort(ServPort)
		//s.SetDumpRouteMap(false)

		// 文件根目录
		filePath := PathRoot + "/files"
		if !gfile.Exists(filePath) {
			if err := gfile.Mkdir(filePath); err != nil {
				panic(err)
			}
		}
		s.AddStaticPath("/files", filePath)
		s.AddStaticPath("/spa", "public")
		// Run Server
		g.Server().Run()
	}()

}
