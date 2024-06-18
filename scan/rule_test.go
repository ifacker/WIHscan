package scan

import (
	"fmt"
	"testing"
	"wihscan/global"
)

func TestRule(t *testing.T) {
	global.RulePath = "../config/wih_rules.yaml"
	body := `xxx`
	a := rule(body)
	fmt.Println(a)
}

func TestRuleLoad(t *testing.T) {
	global.RulePath = "../config/test.yaml"
	RuleLoad()

	// ctm 服了，自己序列化
	// wih := &dataType.WIH{
	// 	Rules: []dataType.Rule{
	// 		dataType.Rule{
	// 			Id:      "1",
	// 			Enabled: true,
	// 			Pattern: "12",
	// 		},
	// 		dataType.Rule{
	// 			Id:      "2",
	// 			Enabled: false,
	// 			Pattern: "123",
	// 		},
	// 	},
	// 	Exclude_rules: []dataType.Exclude_rule{},
	// }
	// a, _ := yaml.Marshal(wih)
	// fmt.Println(string(a))
}
