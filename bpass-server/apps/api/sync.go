package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

type SyncController struct {
	ws *ghttp.WebSocket
}

var ctx = gctx.New()
var ()

// Index 触发页面
// /sync/
func (c *SyncController) Index(r *ghttp.Request) {
	contains, _ := r.Session.Contains("clientId")
	sid, _ := r.Session.Id()
	if !contains {
		_ = r.Session.Set("clientId", sid)
	}
	//_ = c.View.Display("sync.html")
}

// WebSocket 接口
// /sync/web-socket
func (c *SyncController) WebSocket(r *ghttp.Request) {

	// 初始化WebSocket请求
	if ws, err := r.WebSocket(); err == nil {
		c.ws = ws
	} else {
		glog.Error(ctx, err)
		return
	}

	// 初始化时设置用户信息
	clientId, _ := r.Session.Get("clientId")
	if clientId.String() == "" {
		id, _ := r.Session.Id()
		_ = r.Session.Set("clientId", id)
	}
	users.Set(c.ws, clientId)
	names.Add(clientId.String())

	for {
		// 阻塞读取WS数据
		msgType, msg, err := c.ws.ReadMessage()
		if err != nil {
			users.Remove(c.ws)
			names.Remove(clientId.String())
			break
		}

		// 群发同步所有端
		glog.Cat("sync").Print(ctx, "[sync] ", clientId, msg)
		_ = c.writeUsers()
		if msg != nil {
			msgs := "{" +
				"\"clientId\":\"" + clientId.String() + "\"," +
				"\"msg\":\"" + string(msg) + "\"" +
				"}"
			_ = c.writeGroup(msgType, msgs)
		}
	}
}

// 群发消息
func (c *SyncController) writeGroup(msgType int, msg string) error {
	msgs := []byte(msg)
	users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(msgType, []byte(msgs))
		}
	})
	return nil
}

// 向客户端返回用户列表
func (c *SyncController) writeUsers() error {
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
	if err := c.writeGroup(1, msgs); err != nil {
		return err
	}
	return nil
}
