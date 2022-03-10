package fileinfos

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

var ctx = gctx.New()

// Init DataInit 从文件恢复为缓存
func Init(keys ...string) {
	fmt.Println(keys)
	for _, key := range keys {
		data := gfile.GetContents(cacheFile(key))
		gcache.Set(ctx, key, data, 0)
	}
}

// Set DataSet 设置缓存
func Set(key, value string) {

	err := gcache.Set(ctx, key, value, 0)
	if err != nil {
		return
	}
	_ = gfile.PutContents(cacheFile(key), value)
}

// Get DataGet 读取缓存
func Get(key string) string {
	cacheGet, err := gcache.Get(ctx, key)
	if err != nil {
		panic("go cache get fail!!")
	}
	return cacheGet.String()
}

// cacheFile 缓存实例化文件
func cacheFile(key string) string {
	return GetRootPath() + "/tmp/data/" + key + ".txt"
}
