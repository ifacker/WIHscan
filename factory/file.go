package factory

import (
	"wihscan/util"

	"github.com/ifacker/myutil"
)

func factoryFile(filepath string) []string {
	jsUrls := util.ReadFile2Line(filepath)

	newJsUrls := myutil.RemoveDuplicateElement(jsUrls)
	return newJsUrls
}
