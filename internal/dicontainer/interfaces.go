package dicontainer

import (
	"github.com/av-belyakov/simplelogger"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
)

type Logger interface {
	GetChan() <-chan interfaces.Messager
	Send(msgType, message string)
	Close()
}

type Counter interface {
	SendMessage(msgType string, count int)
}

type SimpleLogger interface {
	SetDataBaseInteraction(dbi simplelogger.DataBaseInteractor)
	GetCountFileDescription() int
	GetListTypeFiles() []string
	Write(typeLog, msg string) bool
}

type Configer interface {
	Get() confighandler.Config
	GetCommon() confighandler.Common
	GetNATS() confighandler.CfgNats
	GetKafka() confighandler.CfgKafka
	GetListLogs() []*confighandler.LogSet
	GetStorageDB() confighandler.CfgStorageDB
	GetLogDB() confighandler.CfgWriteLogDB
	GetZabbix() confighandler.ZabbixOptions
}

type NatsConnecter interface {
	GetChannelFromModule() <-chan natsapi.SettingsChanOutput
	GetChannelToModule() chan<- natsapi.SettingsChanInput
}

type KafkaConnecter interface {
	GetChannelFromModule() <-chan kafkaapi.SettingsChanOutput
	GetChannelToModule() chan<- kafkaapi.SettingsChanInput
}

type DB interface {
	GetChannelToModule() chan databasestorageapi.SettingsChanInput
	GetChannelFromModule() chan databasestorageapi.SettingsChanOutput
}

type DbLogger interface {
	Write(msgType, msg string) error
}
