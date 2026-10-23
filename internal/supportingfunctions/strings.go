package supportingfunctions

import (
	"fmt"
	"strings"
)

// GetWhitespace возвращает необходимое количество пробелов
func GetWhitespace(num int) string {
	var str string

	if num == 0 {
		return str
	}

	for range num {
		str += "  "
	}

	return str
}

// JoinRawFieldsToString объединяет список необработанных полей в строку
func JoinRawFieldsToString(list map[string]string, tag, id string) string {
	var str strings.Builder = strings.Builder{}

	for k, v := range list {
		fmt.Fprintf(&str, "\n\t%s %s field: '%s', value: '%s'", tag, id, k, v)
	}

	return str.String()
}

func ToStringBeautifulSlice[T any](num int, l []T) string {
	str := strings.Builder{}
	ws := GetWhitespace(num + 1)

	for k, v := range l {
		fmt.Fprintf(&str, "%s%d. '%v'\n", ws, k+1, v)
	}

	return str.String()
}

func ToStringBeautifulMapSlice(num int, m map[string][]string) string {
	str := strings.Builder{}

	for k, v := range m {
		fmt.Fprintf(&str, "%s'%s'\n", GetWhitespace(num+1), k)
		for key, value := range v {
			fmt.Fprintf(&str, "%s%d. %s\n", GetWhitespace(num+2), key+1, value)
		}
	}

	return str.String()
}
