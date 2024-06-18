package scan

import (
	"fmt"
	"testing"
	"wihscan/global"
)

func TestScan(t *testing.T) {
	global.RulePath = "../config/wih_rules.yaml"
	urljs := "https://health.house.oirlife.com:9002/static/js/chunk-vendors.15e738ce.js"
	result := Scan(urljs)
	fmt.Println(result)
}
