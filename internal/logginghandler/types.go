package logginghandler

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

type LoggingChan struct {
	dataWriter  interfaces.WriterLoggingData
	chanLogging chan interfaces.Messager
}

// MessageLogging содержит информацию используемую при логировании
type MessageLogging struct {
	Message, Type string
}
