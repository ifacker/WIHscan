package factory

import datatype "wihscan/dataType"

// 读取数据，处理数据
func Factory(option *datatype.Option) []string {
	if option.JsonFilePath != "" {
		return factoryJson(option.JsonFilePath)
	}
	if option.FilePath != "" {
		return factoryFile(option.FilePath)
	}
	if option.Url != "" {
		return []string{option.Url}
	}
	return nil
}
