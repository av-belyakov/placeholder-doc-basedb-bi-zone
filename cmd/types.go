package main

import (
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
)

// ApplicationRouter модуль маршрутизации
type ApplicationRouter struct {
	logger         interfaces.Logger
	counter        interfaces.Counter
	chToNatsApi    chan<- natsapi.SettingsChanInput
	chFromNatsApi  <-chan natsapi.SettingsChanOutput
	chToKafkaApi   chan<- kafkaapi.SettingsChanInput
	chFromKafkaApi <-chan kafkaapi.SettingsChanOutput
	chToDBSApi     chan<- databasestorageapi.SettingsChanInput
	chFromDBSApi   <-chan databasestorageapi.SettingsChanOutput
}

// ApplicationRouterSettings настройки модуля маршрутизации
type ApplicationRouterSettings struct {
	ChanToNats    chan<- natsapi.SettingsChanInput
	ChanFromNats  <-chan natsapi.SettingsChanOutput
	ChanToKafka   chan<- kafkaapi.SettingsChanInput
	ChanFromKafka <-chan kafkaapi.SettingsChanOutput
	ChanToDBS     chan<- databasestorageapi.SettingsChanInput
	ChanFromDBS   <-chan databasestorageapi.SettingsChanOutput
}
