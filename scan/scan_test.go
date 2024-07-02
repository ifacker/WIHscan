package scan

import (
	"fmt"
	"testing"
	"wihscan/global"
)

func TestScan(t *testing.T) {
	global.RulePath = "../config/wih_rules.yaml"
	urljs := ""
	result := Scan(urljs)
	fmt.Println(result)
}
