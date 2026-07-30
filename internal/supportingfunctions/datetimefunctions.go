package supportingfunctions

import (
	"fmt"
	"regexp"
	"time"
)

// GetDateTimeFormatRFC3339 конвертирует числовое значение времени Unixtime
// в строку времени в формате RFC3339. Для корректной работы нужна дата в
// формате UnixMilli-секунд (13 символов)
func GetDateTimeFormatRFC3339(dt int64) string {
	return time.UnixMilli(dt).Format(time.RFC3339)
}

// SmartConvertToRFC3339 позволяет гибко конвертировать строку даты различных форматов в строку в формате RFC3339
func SmartConvertToRFC3339(dateStr string) (string, error) {
	var layout string

	// Определяем формат по регулярным выражениям
	switch {
	case regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[Z|\+|-]\d{2}:\d{2}$`).MatchString(dateStr):
		//RFC3339 = "2006-01-02T15:04:05Z07:00"
		layout = time.RFC3339

	case regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+[Z|\+|-]\d{2}:\d{2}$`).MatchString(dateStr):
		//RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		layout = time.RFC3339Nano

	case regexp.MustCompile(`^\w{3}\s\w{3}\s[_|\d]\d\s\d{2}:\d{2}:\d{2}\s\d{4}$`).MatchString(dateStr):
		//ANSIC = "Mon Jan _2 15:04:05 2006"
		layout = time.ANSIC

	case regexp.MustCompile(`^\w{3}\s\w{3}\s[\s|_|\d]\d\s\d{2}:\d{2}:\d{2}\s\w{3}\s\d{4}$`).MatchString(dateStr):
		//UnixDate = "Mon Jan _2 15:04:05 MST 2006"
		layout = time.UnixDate

	case regexp.MustCompile(`^\w{3}\s\w{3}\s[_|\d]\d\s\d{2}:\d{2}:\d{2}\s[-|\+]\d{4}\s\d{4}$`).MatchString(dateStr):
		//RubyDate = "Mon Jan 02 15:04:05 -0700 2006"
		layout = time.RubyDate

	case regexp.MustCompile(`^\d{2}\s\w{3}\s\d{2}\s\d{2}:\d{2}\s\w{3}$`).MatchString(dateStr):
		//RFC822 = "02 Jan 06 15:04 MST"
		layout = time.RFC822

	case regexp.MustCompile(`^\d{2}\s\w{3}\s\d{2}\s\d{2}:\d{2}\s[-|\+]\d{4}$`).MatchString(dateStr):
		//RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		layout = time.RFC822Z

	case regexp.MustCompile(`^\w+,\s\d{2}-\w{3}-\d{2}\s\d{2}:\d{2}:\d{2}\s\w{3}$`).MatchString(dateStr):
		//RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
		layout = time.RFC850

	case regexp.MustCompile(`^\w{3},\s\d{2}\s\w{3}\s\d{4}\s\d{2}:\d{2}:\d{2}\s\w{3}$`).MatchString(dateStr):
		//RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
		layout = time.RFC1123

	case regexp.MustCompile(`^\w{3},\s\d{2}\s\w{3}\s\d{4}\s\d{2}:\d{2}:\d{2}\s[-\+]\d{4}$`).MatchString(dateStr):
		//RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123
		layout = time.RFC1123Z

	case regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d+ [-+]\d{4} [A-Z]+$`).MatchString(dateStr):
		layout = "2006-01-02 15:04:05.999999999 -0700 MST"

	case regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`).MatchString(dateStr):
		layout = "2006-01-02 15:04:05"

	case regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4} \d{2}:\d{2}:\d{2}$`).MatchString(dateStr):
		layout = "02.01.2006 15:04:05"

	case regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(dateStr):
		layout = "2006-01-02"

	default:
		return "", fmt.Errorf("неизвестный формат даты: %s", dateStr)
	}

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return "", err
	}

	return t.Format(time.RFC3339), nil
}
