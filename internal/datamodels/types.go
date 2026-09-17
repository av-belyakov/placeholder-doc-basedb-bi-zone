package datamodels

// SupportingStructureForTags вспомогательный тип используемый для хранения
// объектов типа tags
type SupportingStructureForTags struct {
	listAcceptedFields []string
	tagTmp             BiZoneIRPTag
	tags               []BiZoneIRPTag
}

// SupportingStructureForSnapshots вспомогательный тип используемый для хранения
// объектов типа snapshots
type SupportingStructureForSnapshots struct {
	listAcceptedFields []string
	snapshotTmp        BiZoneIRPSnapshot
	snapshots          []BiZoneIRPSnapshot
}

// AdditionalInformation дополнительная информация добавляемая к информации по кейсам
type AdditionalInformation struct {
	Sensors     []SensorInformation    `json:"@sensorAdditionalInformation"`
	IpAddresses []IpAddressInformation `json:"@ipAddressAdditionalInformation"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	INN         string `json:"inn" bson:"inn"`                 //налоговый идентификатор
	HostId      string `json:"hostId" bson:"hostId"`           //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	OrgName     string `json:"orgName" bson:"orgName"`         //наименование организации
	HomeNet     string `json:"homeNet" bson:"homeNet"`         //перечень домашних сетей
	GeoCode     string `json:"geoCode" bson:"geoCode"`         //географический код страны
	SensorId    string `json:"sensorId" bson:"sensorId"`       //идентификатор сенсора
	SubjectRF   string `json:"subjectRF" bson:"subjectRF"`     //субъект Российской Федерации
	ObjectArea  string `json:"objectArea" bson:"objectArea"`   //сфера деятельности объекта
	FullOrgName string `json:"fullOrgName" bson:"fullOrgName"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressInformation struct {
	Ip          string `json:"ip"`          //ip адрес по которому выполнялся поиск
	City        string `json:"city"`        //город
	Country     string `json:"country"`     //страна
	CountryCode string `json:"countryCode"` //код страны
}

// AdditionalInformationSensors дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationSensors struct {
	Sensors []SensorInformation `json:"@sensor_additional_information"`
}

// AdditionalInformationIpAddress дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationIpAddress struct {
	IpAddresses []IpAddressInformation `json:"@ip_address_additional_information"`
}
