package scan

import (
	"io"
	"net/http"
	"strings"
	datatype "wihscan/dataType"
	"wihscan/util"
)

// 对单个 js 进行扫描，获取敏感信息
func Scan(urljs string) *datatype.ScanResult {
	urljs = strings.TrimSpace(urljs)

	client := util.NewClient()

	req, err := http.NewRequest(http.MethodGet, urljs, nil)
	if err != nil {
		util.ErrPrint(err)
		return nil
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {

		util.ErrPrint(err)
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		util.ErrPrint(err)
		return nil
	}

	// 规则库匹配
	infoMap := rule(string(body))

	result := &datatype.ScanResult{UrlJs: urljs, Info: infoMap}

	return result
}
