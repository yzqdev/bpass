package api

import (
	"b0pass/boot"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"b0pass/library/openurl"
	"b0pass/library/response"
	"github.com/gogf/gf/v2/net/ghttp"
	"strconv"
)

// OpenUrl 打开本地url
func OpenUrl(r *ghttp.Request) {
	getUrl := r.GetQuery("url").String()
	_ = openurl.Open(getUrl)
}

// GlobalData 获取IP地址
func GlobalData(r *ghttp.Request) {

	port := boot.ServPort
	pathRoot := fileinfos.GetRootPath() + "/files/"
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}

	data := map[string]interface{}{
		"ips":      ips,
		"pathRoot": pathRoot,
	}
	response.JSON(r, 0, "ok", data)
}

// GetSubPath GetPathSub 上传目录记忆功能
func GetSubPath(r *ghttp.Request) {
	saveData(r, "path", "data_path")
}

type Text struct {
	Data string `json:"data"`
	Code string `json:"code"`
}

var dataText = "data_text"

// SendTextData   文本内容发送
func SendTextData(r *ghttp.Request) {
	var req *Text
	if err := r.Parse(&req); err != nil {
		response.JSON(r, 400, "二重唱")
		return
	}
	getData := req.Data

	fileinfos.Set(dataText, getData)
	response.JSON(r, 0, "ok", "已发送")

}

// GetTextData 文本内容获取
func GetTextData(r *ghttp.Request) {
	dbData := fileinfos.Get(dataText)
	response.JSON(r, 0, "ok", dbData)
}

// saveData 保存数据到文件
func saveData(r *ghttp.Request, getkey string, dbkey string) {
	var req *Text
	if err := r.Parse(&req); err != nil {
		response.JSON(r, 400, "二重唱")
		return
	}
	getData := req.Data
	getCode := req.Code
	if getCode == "1" {
		fileinfos.Set(dbkey, getData)
	}
	dbData := fileinfos.Get(dbkey)
	//dbData:=time.Now().Format("2006-01-02")
	response.JSON(r, 0, "ok", dbData)
}
