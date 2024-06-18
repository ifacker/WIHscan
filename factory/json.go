package factory

import (
	"wihscan/util"

	"github.com/ifacker/myutil"
)

// 处理 json 格式
func factoryJson(filepath string) []string {
	// jsonStr := util.ReadFile2All(filepath)
	jsonStr := util.ReadFile2Str(filepath)
	jsUrls := util.GetUrl4Js(jsonStr)

	newJsUrls := myutil.RemoveDuplicateElement(jsUrls)
	return newJsUrls
}
