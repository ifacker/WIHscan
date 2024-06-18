package util

func GetUrl4Js(data string) []string {
	results, err := Regex2(data, `(?<="source": ")([^"]+)(?=")`)
	if err != nil {
		ErrPrint(err)
	}
	return results
}
