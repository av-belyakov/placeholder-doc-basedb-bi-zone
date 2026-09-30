package databasestorageapi

import (
	"github.com/elastic/go-elasticsearch/v9"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

type DatabaseStorage struct {
	counter  interfaces.Counter
	logger   interfaces.Logger
	client   *elasticsearch.Client
	settings settingsDatabaseStorage
	chInput  chan SettingsChanInput  //канал для передачи данных в модуль
	chOutput chan SettingsChanOutput //канал для приёма данных из модуля
}

type settingsDatabaseStorage struct {
	storages            map[string]string //хранилища (в elasticsearch - индексы)
	namedb              string            //наименование базы данных
	user                string            //имя пользователя для аутентификации
	passwd              string            //пароль для авторизации
	host                string            //ip адрес или доменное имя
	port                int               //сетевой порт
	maxGetDocumentsSize int               //максимальное количество запрашиваемых документов
}

// SettingsChanInput параметры канала для передачи данных в модуль
type SettingsChanInput struct {
	Data    any
	Section string
	Command string
}

// SettingsChanOutput параметры канала для приёма данных из модуля
type SettingsChanOutput struct {
	Data    []byte
	Id      string
	UUID    string
	Command string
}

type DatabaseStorageOptions func(*DatabaseStorage) error

//****** для объектов типа VerifiedBiZoneAlerts *******

// ResponseVerifiedBiZoneAlerts ответ от базы данных
type ResponseVerifiedBiZoneAlerts struct {
	Options ResponseVerifiedBiZoneAlertsOptions `json:"hits"`
}

// ResponseVerifiedBiZoneAlertsOptions опции ответа
type ResponseVerifiedBiZoneAlertsOptions struct {
	Total    OptionsTotal                 `json:"total"`
	Hits     []PatternVerifiedBiZoneAlert `json:"hits"`
	MaxScore float64                      `json:"max_score"`
}

// PatternVerifiedBiZoneAlert шаблон
type PatternVerifiedBiZoneAlert struct {
	Source datamodels.VerifiedBiZoneIRPAlert `json:"_source"`
	ServiseOption
}

//****** для объектов типа AdditionalInformation *******

// ResponseTemplateAdditionalInformation шаблон для дополнительной информации
type ResponseTemplateAdditionalInformation struct {
	Options TemplateAdditionalInformationOptions `json:"hits"`
}

// TemplateAdditionalInformationOptions шаблон опций
type TemplateAdditionalInformationOptions struct {
	Total    OptionsTotal                    `json:"total"`
	Hits     []TemplateAdditionalInformation `json:"hits"`
	MaxScore float64                         `json:"max_score"`
}

// TemplateAdditionalInformation шаблон документа
type TemplateAdditionalInformation struct {
	Source datamodels.AdditionalInformation `json:"_source"`
	ServiseOption
}

// OptionsTotal опции в результате поиска
type OptionsTotal struct {
	Relation string `json:"relation"` //отношение (==, >, <)
	Value    int    `json:"value"`    //количество найденных значений
}

// ServiseOption сервисные опции
type ServiseOption struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}
