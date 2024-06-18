package dataType

// WIH 规则配置
type WIH struct {
	Rules         []Rule         `yaml:"rules"`
	Exclude_rules []Exclude_rule `yaml:"exclude_rules"`
}

// 规则
type Rule struct {
	Id      string `yaml:"id"`
	Enabled bool   `yaml:"enabled"`
	Pattern string `yaml:"pattern"`
}

// 排除的规则
// 排除规则， 支持字段 id, content, target , source 逻辑为 and ，如果是正则匹配，需要使用 regex: 开头
type Exclude_rule struct {
	Name       string `yaml:"name"`
	Id         string `yaml:"id"`
	Target     string `yaml:"target"`
	Enabled    bool   `yaml:"enabled"`
	Content    string `yaml:"content"`
	Source     string `yaml:"source"`
	Source_tag string `yaml:"source_tag"`
	Regex      string `yaml:"regex"`
}
