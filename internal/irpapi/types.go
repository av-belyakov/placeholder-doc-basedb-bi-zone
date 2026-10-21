package irpapi

import (
	"encoding/json"
)

type IrpApi struct {
	settings ApiSettings
}

// settings настройки модуля
type ApiSettings struct {
	apiKey string
	host   string
	port   int
}

// options функциональные опции
type options func(*IrpApi) error

type InputReguest struct {
	Data    []byte
	Command string
	RootId  string
}

// RequestCommand структура с командами для обработки модулем
type RequestCommand struct {
	ByteData       json.RawMessage `json:"byte_data"`           //набор данных в бинарном виде которые обрабатываются отдельно
	Service        string          `json:"service"`             //наименование сервиса
	Command        string          `json:"command"`             //команда
	RootId         string          `json:"root_id"`             //основной id, как правило это rootId case или alert
	CaseId         string          `json:"case_id"`             //id кейса
	Value          string          `json:"value"`               //устанавливаемое значение
	Username       string          `json:"username"`            //имя пользователя, необходим если нужно указать пользователя выполнившего действие
	FieldName      string          `json:"field_name"`          //некое ключевое поле
	RegionalObject string          `json:"for_regional_object"` //для кого предназначена команда
}

// Querys перечень запросов к TheHive
type Querys struct {
	Query []Query `json:"query"`
}

// Query структура запроса к TheHive
type Query struct {
	ExtraData []string `json:"extraData,omitempty"`
	Name      string   `json:"_name,omitempty"`
	IDOrName  string   `json:"idOrName,omitempty"`
	From      int64    `json:"from"`
	To        int      `json:"to,omitempty"`
}
