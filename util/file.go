package util

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed lib/rules.yaml
var libRules embed.FS

// 读取文件，按行返回 []string
func ReadFile2Line(filepath string) []string {
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("打开文件时出错:", err)
		return nil
	}
	defer file.Close()

	// 创建一个新的 Scanner 对象
	scanner := bufio.NewScanner(file)

	results := []string{}

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, strings.TrimSpace(line))
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
		return nil
	}

	return results
}

// 读取文件，读取全部内容
func ReadFile2All(filepath string) string {
	resultAll := ""
	results := ReadFile2Line(filepath)
	for _, result := range results {
		resultAll += fmt.Sprintf("%s\n", result)
	}
	return resultAll
}

// 读取整个文件并返回 byte
func ReadFile2Byte(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		ErrPrint(err)
		return nil
	}
	return content
}

// 读取整个文件并返回 string
func ReadFile2Str(filepath string) string {
	return string(ReadFile2Byte(filepath))
}

// 写入文件
func WriteFile(filepath, content string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		ErrPrint(err)
		return
	}
	file.Write([]byte(content))
}

// 初始化 规则配置文件
func InitRule() {
	// 判断配置文件是否存在
	_, err := os.Lstat("config/rules.yaml")
	if !os.IsNotExist(err) {
		return
	}
	_, err = os.Lstat("config/")
	if os.IsNotExist(err) {
		// 创建多级文件夹
		err := os.MkdirAll("config/", 0755)
		if err != nil {
			ErrPrint(err)
			return
		}
	}
	content, err := libRules.ReadFile("lib/rules.yaml")
	if err != nil {
		ErrPrint(err)
	}
	WriteFile("config/rules.yaml", string(content))
}
