package util

import (
	"net/http"
	"wihscan/global"
)

// 一统 http.client
func NewClient() *http.Client {
	client := &http.Client{
		Transport: global.Tr,
		Timeout:   global.TimeOut,
	}
	return client
}
