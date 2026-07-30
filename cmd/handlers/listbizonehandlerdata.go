package handlers

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

// NewListBiZoneHandlerData начальный обработчик событий топика 'data.*'
func NewListBiZoneHandlerData(data *datamodels.BiZoneIRPData) map[string][]func(any) {
	return map[string][]func(any){
		"data._id":           {data.SetAnyID},
		"data.title":         {data.SetAnyTitle},
		"data.sensor":        {data.SetAnySensor},
		"data.ip_home":       {data.SetAnyIPHome},
		"data.ip_exter":      {data.SetAnyIPExter},
		"data.url____ftp":    {data.SetAnyURLFTP},
		"data.url___http":    {data.SetAnyURLHTTP},
		"data.event_type":    {data.SetAnyEventType},
		"data.url_arkime":    {data.SetAnyURLArkime},
		"data.all____ip_ext": {data.SetAnyAllIPExt},
		"data.response_team": {data.SetAnyResponseTeam},
		//ниже работа со срезам содержащими простые типы
		"data.snort_sid":    {data.SetAnySnortSid},
		"data.all_sensors":  {data.SetAnyAllSensor},
		"data.all__ip_home": {data.SetAnyAllIPHome},
	}
}
