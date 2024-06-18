package options

import (
	"time"
	datatype "wihscan/dataType"
	"wihscan/global"
	"wihscan/util"

	"github.com/gookit/color"
	"github.com/ifacker/myutil"
	"github.com/projectdiscovery/goflags"
)

var option = &datatype.Option{}

func init() {
	// 初始化配置文件
	util.InitRule()

	flagset := goflags.NewFlagSet()

	// 区分大小写
	flagset.CaseSensitive = false

	flagset.CreateGroup("option", "必选参数（这里 -u、-f、-fj 三选其一）",
		flagset.StringVarP(&option.Url, "url", "u", "", "需要扫描的 url"),
		flagset.StringVarP(&option.FilePath, "filePath", "f", "", "要扫描的文件路径"),
		flagset.StringVarP(&option.JsonFilePath, "jsonPath", "fj", "", "要扫描的json文件的路径"),
	)

	flagset.CreateGroup("other", "可选参数",
		flagset.StringVarP(&option.OutputFilePath, "output", "o", "result.txt", "输出路径，默认当前目录的 result.txt 文件中"),
		flagset.IntVarP(&option.Thread, "thread", "t", 200, "设置并发量"),
		flagset.StringVarP(&option.Proxy, "proxy", "p", "", "代理地址，如：-p http://localhost:8080 或 -p socks5://localhost:1080"),
		flagset.IntVarP(&option.TimeOut, "timeout", "", 60, "设置超时时长，单位：秒"),
		flagset.StringVarP(&option.RuleConfigPath, "rule", "", "config/rules.yaml", "指定规则配置文件路径"),
		flagset.BoolVarP(&option.Debug, "debug", "", false, "debug 模式，开启错误提示（方便开发使用）"),
	)
	flagset.Parse()
}

func Options() *datatype.Option {
	// 初始化代理
	if option.Proxy != "" {
		myutil.InitProxy(global.Tr, option.Proxy)
	}

	if option.FilePath == "" && option.JsonFilePath == "" && option.Url == "" {
		color.Errorln("请用 -u / -f / -fj 输入需要扫描的 URL 或 文件")
		return nil
	}

	global.TimeOut = time.Duration(option.TimeOut) * time.Second
	global.RulePath = option.RuleConfigPath

	return option
}
