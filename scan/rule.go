package scan

import (
	"fmt"
	datatype "wihscan/dataType"
	"wihscan/global"
	"wihscan/util"

	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

// 对规则进行匹配
func rule(body string) map[string]string {
	resultMap := make(map[string]string)

	if global.RuloWIH == nil {
		return nil
	}

	// 处理 rule
	for _, rule := range global.RuloWIH.Rules {
		if !rule.Enabled {
			continue
		}
		if rule.Pattern == "" {
			continue
		}
		results, err := util.Regex2(body, rule.Pattern)
		if err != nil {
			color.Errorf("%s 此规则异常，将直接跳过！\n%s\n", rule.Id, err)
			continue
		}
		if results == nil {
			continue
		}
		for i, result := range results {
			if i == 0 {
				resultMap[rule.Id] = result
			} else {
				resultMap[fmt.Sprintf("%s_%d", rule.Id, i)] = result
			}
		}
	}
	return resultMap
}

// 加载规则
func RuleLoad() {
	global.RuloWIH = &datatype.WIH{}

	configFile := util.ReadFile2Byte(global.RulePath)

	err := yaml.Unmarshal(configFile, global.RuloWIH)
	if err != nil {
		util.ErrPrint(err)
		return
	}

	for _, rule := range global.RuloWIH.Rules {
		if rule.Pattern == "" {
			color.Warnf("%s 此规则跳过！\n", rule.Id)
		}
	}
	fmt.Println()

}
