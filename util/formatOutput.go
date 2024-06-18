package util

import (
	"fmt"
	"wihscan/dataType"

	"github.com/gookit/color"
)

// 格式化打印输出
func FormatOutput(resultScan *dataType.ScanResult) {
	if resultScan == nil {
		return
	}
	if len(resultScan.Info) <= 0 {
		return
	}

	format := fmt.Sprintf("[+] %s\n", resultScan.UrlJs)
	for k, v := range resultScan.Info {
		format += color.C256(51).Sprintf("    %s: %s\n", k, v)
	}

	color.C256(46).Println(format)
}

// 格式化写入文件
func FormatOutputWrite(resultScan *dataType.ScanResult, writePath string) {
	if resultScan == nil {
		return
	}
	if len(resultScan.Info) <= 0 {
		return
	}

	format := fmt.Sprintf("[+] %s\n", resultScan.UrlJs)
	for k, v := range resultScan.Info {
		format += fmt.Sprintf("    %s: %s\n", k, v)
	}

	WriteFile(writePath, format+"\n")
}
