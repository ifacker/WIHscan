package dataType

type Option struct {
	Url            string // 单个 URL
	FilePath       string // 文件路径
	JsonFilePath   string // json 文件路径
	Proxy          string // 代理地址
	Thread         int    // 并发量
	Debug          bool   // debug 模式
	TimeOut        int    // 超时时间
	OutputFilePath string // 输出路径
	RuleConfigPath string // 制定规则配置文件
}
