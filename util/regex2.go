package util

import "github.com/dlclark/regexp2"

// 传入要匹配的内容 body 和正则表达式 regex
func Regex2(body, regex string) ([]string, error) {
	re, err := regexp2.Compile(regex, 0)
	if err != nil {
		return nil, err
	}
	regexResult := regexp2FindAllString(re, body)
	return regexResult, nil
}

// 正则表达式 全局查找 支持最新规则 （传入 re 参数和将要筛选的内容）
func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
