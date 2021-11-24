package api

import (
	"b0pass/library/fileinfos"
	"b0pass/library/response"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/gookit/color"
	"io"
	"log"
	"os"
	"reflect"
)

// Upload 执行文件上传处理
func Upload(r *ghttp.Request) {
	//if err := r.ParseMultipartForm(32); err != nil {
	//	response.JSON(r, 201, err.Error())
	//}
	file := r.GetUploadFiles("file")

	//defer func() { _ = file.close() }()

	// Get path
	pathSub := r.GetString("path")
	fileinfos.Set("data_path", pathSub)
	// Save path
	savePath := fileinfos.GetRootPath() + "/files/" + pathSub + "/"
	log.Println(savePath)
	// Upload file
	_, err := file.Save(savePath)
	if err != nil {
		response.JSON(r, 201, "ok", false)
		color.Redln(err)
		return
	}

	response.JSON(r, 0, "ok", true)

}
func MultiUpload(r *ghttp.Request) {
	file := r.GetUploadFiles("file")
	fmt.Println(file)
	savePath := fileinfos.GetRootPath() + "/files/" + " name"
	_, err := file.Save(savePath)
	if err != nil {
		color.Redln("save fail")
		return
	}
	// Upload file
	fmt.Println(reflect.TypeOf(file))
}

// Upload 以小内存上传大文件
func Uploadx(r *ghttp.Request) {
	//Multipart Pipe
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		response.JSON(r, 201, err.Error())
	}
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() { _ = f.Close() }()
		name := gfile.Basename(h.Filename)

		//写入文件
		dst, err := os.OpenFile(
			fileinfos.GetRootPath()+"/files/"+name,
			os.O_WRONLY|os.O_CREATE, 0666,
		)
		defer func() { _ = dst.Close() }()
		if err != nil {
			response.JSON(r, 201, err.Error())
		}
		if _, err := io.Copy(dst, f); err != nil {
			response.JSON(r, 201, err.Error())
		}

		response.JSON(r, 0, "ok", name)
	} else {
		response.JSON(r, 201, e.Error())
	}
}

// Lists
func Lists(r *ghttp.Request) {
	fp := fileinfos.GetRootPath() + "/files/*"
	var ret []map[string]string
	ret = fileinfos.ListDirData(fp, "files")
	response.JSON(r, 0, "ok", ret)
}

// Delete
func Delete(r *ghttp.Request) {
	f := r.Get("f")
	fp := fileinfos.GetRootPath()
	filePath := fp + "/files/" + gconv.String(f)
	err := os.RemoveAll(filePath)
	if err != nil {
		response.JSON(r, 0, "failed", filePath)
		return
	}
	response.JSON(r, 0, "ok", filePath)
}

func Dump(r *ghttp.Request) {
	filePath := os.Args[0]
	response.JSON(r, 0, "ok", filePath)
}

// UploadShow 展示文件上传页面
func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>上传文件</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/api/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}
