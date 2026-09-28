package documentsgenerator

import (
	"fmt"
	"testing"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
	"github.com/stretchr/testify/assert"
)

func TestTimeDecode(t *testing.T) {
	var timeDecodeString string = "2026-09-17T05:52:00Z"

	t.Run("Тест 1. Декодирование строки времени в формате RFC 3339", func(t *testing.T) {
		timeRFC3339, err := supportingfunctions.SmartConvertToRFC3339("2026-09-17T10:58:37+03:00")
		assert.NoError(t, err)

		fmt.Println("time format RFC 3339:", timeRFC3339)
	})

	t.Run("Тест 2. Декодирование строки времени в другом формате", func(t *testing.T) {
		someTime, err := supportingfunctions.SmartConvertToRFC3339(timeDecodeString)
		assert.NoError(t, err)

		fmt.Println("some time format:", someTime)
	})

	t.Run("Test decode", func(t *testing.T) {
		timeStr, err := time.Parse(time.RFC3339, timeDecodeString)
		assert.NoError(t, err)

		fmt.Println("Time decode:", timeStr)
	})
}
