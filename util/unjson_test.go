package util

import (
	"fmt"
	"testing"
)

func TestUnjson(t *testing.T) {
	str := ``

	data := Unjson(str)
	fmt.Println(data)
}
