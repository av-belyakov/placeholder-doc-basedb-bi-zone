package dicontainer

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

// DiContainer DI контейнер
type DiContainer struct {
	logger       Logger
	counter      Counter
	configer     Configer
	simpleLogger SimpleLogger

	db       DB
	dbLogger DbLogger
	nats     NatsConnecter
	kafka    KafkaConnecter

	ch      chan interfaces.Messager
	rootDir string
}
