package supportingfunctions_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
	"github.com/stretchr/testify/assert"
)

func TestSmartConvertToRFC3339(t *testing.T) {
	listTesting := map[string]struct {
		timeFormat    string
		statusConvert bool
	}{
		time.Now().Format(time.RFC3339):           {"RFC3339", true},
		time.Now().Format(time.RFC3339Nano):       {"RFC3339Nano", true},
		time.Now().Format(time.ANSIC):             {"ANSIC", true},
		time.Now().Format(time.UnixDate):          {"UnixDate", true},
		time.Now().Format(time.RubyDate):          {"RubyDate", true},
		time.Now().Format(time.RFC822):            {"RFC822", true},
		time.Now().Format(time.RFC822Z):           {"RFC822Z", true},
		time.Now().Format(time.RFC850):            {"RFC850", true},
		time.Now().Format(time.RFC1123):           {"RFC1123", true},
		time.Now().Format(time.RFC1123Z):          {"RFC1123Z", true},
		"2026-01-02 23:14:15.999999999 -0700 MST": {"some pattern", true},
		"2026-05-12 20:24:35":                     {"some pattern", true},
		"12.11.2026 23:14:05":                     {"some pattern", true},
		"2026-11-02":                              {"some pattern", true},
		"20-23-44 42:21;30":                       {"some bad pattern", false},
	}

	var count int
	for k, v := range listTesting {
		count++
		t.Run(fmt.Sprintf("Test %d. '%s'", count, k), func(t *testing.T) {
			strRFC3339, err := supportingfunctions.SmartConvertToRFC3339(k)
			if err == nil {
				assert.True(t, v.statusConvert)
			} else {
				fmt.Println("Error:", err)
				assert.False(t, v.statusConvert)
			}

			fmt.Printf("'%s' to format RFC3339 '%s'\n", v.timeFormat, strRFC3339)
		})
	}
}
