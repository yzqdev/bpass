package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"

)

func delete1( ){
	err := os.Remove("D:/GolandProjects/bpass/bpass-server/bin/README.md")

	if err != nil {
		fmt.Println(err)
		color.Redln("hhhh错误")
		return
	}
}
func main() {
	delete1()
}