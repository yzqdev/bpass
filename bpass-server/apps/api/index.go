package api

import (
	"b0pass/boot"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"b0pass/library/response"
	"github.com/gogf/gf/v2/net/ghttp"

	"strconv"
	"time"
)

func Index(r *ghttp.Request) {
	data := map[string]interface{}{
		"times": time.Now().Unix(),
	}
	response.JSON(r, 200, "success", data)
}
func Ips(r *ghttp.Request) {
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}
	data := map[string]interface{}{
		"ips": ips,
	}
	response.JSON(r, 200, "成功", data)

}
func FileLists(r *ghttp.Request) {
	// Ip lists
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}

	// path
	pathRoot := fileinfos.GetRootPath() + "/files/"

	// file lists
	fprPath := r.GetQuery("path").String()
	var fpPath string
	if fprPath != "" {
		fpPath = "/files" + fprPath + "/*"
	} else {
		fpPath = "/files/*"
	}
	fp := fileinfos.GetRootPath() + fpPath
	flists := fileinfos.ListDirData(fp, fprPath)

	// views
	data := map[string]interface{}{
		"ips":      ips,
		"pathRoot": pathRoot,
		"fileList": flists,
	}
	response.JSON(r, 200, "成功", data)

}
