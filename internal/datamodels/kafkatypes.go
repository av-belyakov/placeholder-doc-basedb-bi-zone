package datamodels

// IrpIncidentType заголовок инцидента получаемого через Kafka
// Наиболее интересно поле 'Type' которое содержит один из двух типов описания инцидента:
// - soar.core.incident.created.v1 - для новых инцидентов
// - soar.core.incident.changed.v1 - для уже существующих инцидентов, данные в которых изменились
type IrpIncidentType struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
