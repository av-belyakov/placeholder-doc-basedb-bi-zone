package supportingfunctions

import (
	"regexp"
	"strings"
)

// GetContentFromSnortRule возвращает значения поля 'content' из правила Snort
func GetContentFromSnortRule(snortRule string) ([]string, error) {
	suffies := []string{
		"content:\"",
		"content: \"",
	}

	result := []string(nil)
	rgx, err := regexp.Compile(`content:\s?"([^"]*)";`)
	if err != nil {
		return result, err
	}

	matchFound := rgx.FindAllString(snortRule, -1)
	for _, v := range matchFound {
		for _, suffix := range suffies {
			str := strings.TrimSuffix(strings.TrimPrefix(v, suffix), "\";")
			result = append(result, strings.ReplaceAll(str, "\"", ""))
		}
	}

	return result, nil
}
