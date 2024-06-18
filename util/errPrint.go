package util

import (
	"wihscan/global"

	"github.com/gookit/color"
)

func ErrPrint(err error) {
	if !global.Debug {
		return
	}
	color.Errorln(err)
}
