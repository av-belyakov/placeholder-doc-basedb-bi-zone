package dicontainer

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

// DiContainer DI контейнер
type DiContainer struct {
	logger       Logger
	counter      Counter
	configer     Configer
	simpleLogger SimpleLogger

	dbLogger       DbLogger
	dbConnecter    DBConnecter
	natsConnecter  NatsConnecter
	kafkaConnecter KafkaConnecter

	rootDir string
	ch      chan interfaces.Messager
}
