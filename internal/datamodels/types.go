package datamodels

// SupportingStructureForSnapshots вспомогательный тип используемый для хранения
// объектов типа snapshots
type SupportingStructureForSnapshots struct {
	listAcceptedFields []string
	snapshotTmp        BiZoneIRPSnapshot
	snapshots          []BiZoneIRPSnapshot
}

// AdditionalInformation дополнительная информация добавляемая к информации по кейсам
type AdditionalInformation struct {
	AdditionalInformationSensors
	AdditionalInformationIpAddress
	SpecialUUID string `json:"@special_uuid"`
}

// AdditionalInformationSensors дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationSensors struct {
	Sensors []SensorInformation `json:"@sensor_additional_information"`
}

// AdditionalInformationIpAddress дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationIpAddress struct {
	IpAddresses []IpAddressInformation `json:"@ip_addresses_additional_information"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	INN         string `json:"inn"`           //налоговый идентификатор
	HostId      string `json:"host_id"`       //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	OrgName     string `json:"org_name"`      //наименование организации
	HomeNet     string `json:"home_net"`      //перечень домашних сетей
	GeoCode     string `json:"geo_code"`      //географический код страны
	SensorId    string `json:"sensor_id"`     //идентификатор сенсора
	SubjectRF   string `json:"subject_rf"`    //субъект Российской Федерации
	ObjectArea  string `json:"object_area"`   //сфера деятельности объекта
	FullOrgName string `json:"full_org_name"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressInformation struct {
	Ip          string `json:"ip"`           //ip адрес по которому выполнялся поиск
	City        string `json:"city"`         //город
	Country     string `json:"country"`      //страна
	CountryCode string `json:"country_code"` //код страны
}
