package api

import (
	"b0pass/library/response"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type SyncController struct {
	gmvc.Controller
	ws *ghttp.WebSocket
}

var (
//users = gmap.New()
//names = gset.NewStrSet()
)

// Index 触发页面
// /sync/
func SyncIndex(r *ghttp.Request) {
	if !r.Session.Contains("clientId") {
		_ = r.Session.Set("clientId", r.Session.Id())
	}
	response.JSON(r, 200, "成功", "nodata")

}

// WebSocket 接口
// /sync/web-socket
func WebSocket(r *ghttp.Request) {

	// 初始化WebSocket请求
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}

	// 初始化时设置用户信息
	clientId := r.Session.GetString("clientId")
	if clientId == "" {
		_ = r.Session.Set("clientId", r.Session.Id())
	}
	users.Set(ws, clientId)
	names.Add(clientId)

	for {
		// 阻塞读取WS数据
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			users.Remove(ws)
			names.Remove(clientId)
			break
		}

		// 群发同步所有端
		glog.Cat("sync").Println("[sync] ", clientId, msg)
		_ = writeUsers()
		if msg != nil {
			msgs := "{" +
				"\"clientId\":\"" + clientId + "\"," +
				"\"msg\":\"" + string(msg) + "\"" +
				"}"
			_ = writeGroup(msgType, msgs)
		}
	}
}

// 群发消息
func writeGroup(msgType int, msg string) error {
	msgs := []byte(msg)
	users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(msgType, []byte(msgs))
		}
	})
	return nil
}

// 向客户端返回用户列表
func writeUsers() error {
	nameStr := ""
	names.Iterator(func(v string) bool {
		if nameStr == "" {
			nameStr += v
		} else {
			nameStr += "," + v
		}
		return true
	})
	msgs := "{" +
		"\"clientId\":\"0\"," +
		"\"msg\":\"" + nameStr + "\"" +
		"}"
	if err := writeGroup(1, msgs); err != nil {
		return err
	}
	return nil
}
