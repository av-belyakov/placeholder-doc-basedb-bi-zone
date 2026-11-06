package dicontainer

// DiContainer DI контейнер
type DiContainer struct {
	logger       Logger
	configer     Configer
	simpleLogger SimpleLogger

	dbLogger       DbLogger
	dbConnecter    DBConnecter
	irpConnecter   IRPConnecter
	natsConnecter  NatsConnecter
	kafkaConnecter KafkaConnecter

	rootDir string
}
