package decoderjsondocuments

import (
	"encoding/json"
	"errors"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// New конструктор нового обработчика JSON сообщения
func New(logger interfaces.Logger) *DecoderJsonDocuments {
	return &DecoderJsonDocuments{logger: logger}
}

// Start инициализация обработки
func (s *DecoderJsonDocuments) Start(b []byte) chan interfaces.CustomJsonDecoder {
	chanInput := make(chan interfaces.CustomJsonDecoder)

	go func() {
		//для карт
		listMap := map[string]any{}
		if err := json.Unmarshal(b, &listMap); err == nil {
			if len(listMap) == 0 {
				s.logger.Send("error", supportingfunctions.CustomError(errors.New("error decoding the json message, it may be empty")).Error())

				return
			}

			_ = processingReflectMap(chanInput, listMap, "")
		} else {
			// для срезов
			listSlice := []any{}
			if err = json.Unmarshal(b, &listSlice); err != nil {
				s.logger.Send("error", supportingfunctions.CustomError(err).Error())

				return
			}

			if len(listSlice) == 0 {
				s.logger.Send("error", supportingfunctions.CustomError(errors.New("error decoding the json message, it may be empty")).Error())

				return
			}

			_ = processingReflectSlice(chanInput, listSlice, "")
		}

		close(chanInput)
	}()

	return chanInput
}
