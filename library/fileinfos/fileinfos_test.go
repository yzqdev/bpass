package fileinfos

import (
	"github.com/gookit/color"
	"strings"
	"testing"
)

func TestIfImage(t *testing.T) {
	tmp := IfImage("212/wqewqe/sadsad.png")
	t.Log(tmp)

	tmp = IfImage("212/wqewqe/sadsad.pdf")
	t.Log(tmp)
	color.Cyan.Print(strings.Index(".png",".png"))
}
