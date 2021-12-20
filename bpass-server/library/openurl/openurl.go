package openurl

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// commands 执行程序
var commands = map[string]string{
	"windows": "cmd /c start ",
	"darwin":  "open ",
	"linux":   "xdg-open ", //eog -w
}

// Open 打开浏览器
func Open(uri string) error {
	//runtime.GOOS
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("不存在的操作系统-> %s platform", runtime.GOOS)
	}
	//exec.Command
	run = run + uri
	cmds := strings.Split(run, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	//cmd.Start
	fmt.Println("[CommandAs]", cmds)
	return cmd.Start()
}
