package fileinfos

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// GetBinPath 获取当前可执行文件的路径
func GetBinPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Replace(dir, "\\", "/", -1)
	str = strings.TrimRight(str, "/")
	return str
}

// GetCodePath 获取当前代码文件路径
func GetCodePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Can not get current file info")
	}
	return filepath.Dir(file)
}

// 获取根目录
func GetRootPath() string {
	var fp string
	//fmt.Println("os.Args >>>>> ",os.Args[0][0],os.Args)
	/*if os.Args[0][0]==47 {//exe 47==/
		fp=strings.Replace(GetCodePath(),"/library/fileinfos","",-1)
	}else{
		fp=GetBinPath()
	}*/
	fp = GetBinPath()
	return fp
}

// 根据文件名判断是否是图片
func IfImage(f string) bool {
	var imgs = []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".ico"}
	imgf := false
	for _, imge := range imgs {
		if strings.Index(strings.ToLower(f), imge) > 0 {
			imgf = true
		}
	}
	return imgf
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize  int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
func GetSize(fileBytes uint64) string {
	var (
		units []string
		size  string
		i     int
	)
	units = []string{"B", "K", "M", "G", "T", "P"}
	i = 0
	for {
		i++
		fileBytes = fileBytes / 1024
		if fileBytes < 1024 {
			size = fmt.Sprintf("%v", fileBytes) + units[i]
			break
		}
	}
	return size
}

// ListDirData List Dir Data
func ListDirData(fp, fpSub string) []map[string]string {
	files, _ := filepath.Glob(fp)
	var ret []map[string]string
	var indexs = 0
	for _, file := range files {
		fileInfo, _ := os.Stat(file)
		//filename
		mfile := filepath.Base(file)
		if string(mfile[0]) == "." {
			continue
		}
		//filetype
		mtype := "file"
		if IfImage(mfile) {
			mtype = "img"
		}
		//fileext
		mext := strings.ToUpper(path.Ext(mfile))
		if fileInfo.IsDir() {
			mext = "dir"
			mtype = "dir"
		}
		if mext == "" {
			mext = "file"
			mtype = "file"
		}
		//index
		indexs++
		//map
		m := make(map[string]string)
		m["name"] = mfile
		m["ext"] = mext
		m["size"] = strconv.Itoa(int(fileInfo.Size()))
		m["sizes"] = FormatFileSize(fileInfo.Size())
		m["date"] = fileInfo.ModTime().Format("2006-01-02 15:04:05")
		m["path"] = fpSub + "/" + mfile
		m["type"] = mtype
		m["indexs"] = strconv.Itoa(indexs)
		ret = append(ret, m)
	}
	return ret
}
