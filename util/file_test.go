package util

import (
	"fmt"
	"testing"
	"wihscan/global"
)

func TestReadFile2All(t *testing.T) {
	result := ReadFile2All("unjson.go")
	fmt.Println(result)
}

func TestInitRule(t *testing.T) {
	global.Debug = true
	InitRule()
}
