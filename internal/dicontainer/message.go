package dicontainer

import (
	"context"
	"log"
	"os"

	"github.com/av-belyakov/simplelogger"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/elasticsearchapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/logginghandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// Configer чтения конфигурационного файла
func (d *DiContainer) Configer() Configer {
	if d.configer == nil {
		rootPath, err := supportingfunctions.GetPathRoot(d.rootDir)
		if err != nil {
			log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
		}

		cfg, err := confighandler.New(rootPath)
		if err != nil {
			log.Fatal("error module 'confighandler':", err)
		}

		d.configer = cfg
	}

	return d.configer
}

// SimpleLogger простое логирование с помощью стороннего пакета
func (d *DiContainer) SimpleLogger(ctx context.Context) SimpleLogger {
	if d.simpleLogger == nil {
		listLog := make([]simplelogger.OptionsManager, 0, len(d.Configer().GetListLogs()))
		for _, v := range d.Configer().GetListLogs() {
			listLog = append(listLog, v)
		}

		opts := simplelogger.CreateOptions(listLog...)
		simpleLogger, err := simplelogger.NewSimpleLogger(ctx, d.rootDir, opts)
		if err != nil {
			log.Fatal("error module 'simplelogger':", err)
		}

		d.simpleLogger = simpleLogger

		//подключение логирования в БД
		simpleLogger.SetDataBaseInteraction(d.DbLogger())
	}

	return d.simpleLogger
}

// Logger основное логирование
func (d *DiContainer) Logger(ctx context.Context) Logger {
	if d.logger == nil {
		logger := logginghandler.New(d.SimpleLogger(ctx), d.ch)
		logger.Start(ctx)

		d.logger = logger
	}

	return d.logger
}

// Counter счетчик сообщений
func (d *DiContainer) Counter(ctx context.Context) Counter {
	if d.counter == nil {
		counter := countermessage.New(d.ch)
		counter.Start(ctx)

		d.counter = counter
	}

	return d.counter
}

// DbLogger запись логов в БД
func (d *DiContainer) DbLogger() DbLogger {
	if d.dbLogger == nil {
		var nameRegionalObject = "gcm"
		if os.Getenv("GO_PHMISP_MAIN") == "development" {
			nameRegionalObject = "gcm-test"
		}

		conn, err := elasticsearchapi.NewElasticsearchConnect(elasticsearchapi.Settings{
			Port:               d.Configer().GetLogDB().Port,
			Host:               d.Configer().GetLogDB().Host,
			User:               d.Configer().GetLogDB().User,
			Passwd:             d.Configer().GetLogDB().Passwd,
			IndexDB:            d.Configer().GetLogDB().StorageNameDB,
			NameRegionalObject: nameRegionalObject,
		})
		if err != nil {
			log.Fatal("error module 'elasticsearchapi':", err)
		}

		d.dbLogger = conn
	}

	return d.dbLogger
}

// DB подключение к основной БД
func (d *DiContainer) DB(ctx context.Context) DB {
	if d.db == nil {
		apiDBS, err := databasestorageapi.New(
			d.Logger(ctx),
			d.Counter(ctx),
			databasestorageapi.WithHost(d.Configer().GetStorageDB().Host),
			databasestorageapi.WithPort(d.Configer().GetStorageDB().Port),
			databasestorageapi.WithNameDB(d.Configer().GetStorageDB().NameDB),
			databasestorageapi.WithUser(d.Configer().GetStorageDB().User),
			databasestorageapi.WithPasswd(d.Configer().GetStorageDB().Passwd),
			databasestorageapi.WithStorage(d.Configer().GetStorageDB().Storage),
			databasestorageapi.WithMaxGetDocumentSize(15))
		if err != nil {
			log.Fatal("error initialization module 'db':", err)
		}

		if err := apiDBS.Start(ctx); err != nil {
			log.Fatal("error start module 'db':", err)
		}

		d.db = apiDBS
	}

	return d.db
}

// NatsConnecter подключение к NATS
func (d *DiContainer) NatsConnecter(ctx context.Context) NatsConnecter {
	if d.nats == nil {
		apiNats, err := natsapi.New(
			d.Logger(ctx),
			d.Counter(ctx),
			natsapi.WithHost(d.Configer().GetNATS().Host),
			natsapi.WithPort(d.Configer().GetNATS().Port),
			natsapi.WithCacheTTL(d.Configer().GetNATS().CacheTTL),
			natsapi.WithRequests(d.Configer().GetNATS().EnrichingQueries),
			natsapi.WithSubscriptions(d.Configer().GetNATS().Subscriptions),
			natsapi.WithNameRegionalObject(d.Configer().GetCommon().RegionalObject),
		)
		if err != nil {
			log.Fatal("error initialization module 'natsapi':", err)
		}

		if err = apiNats.Start(ctx); err != nil {
			log.Fatal("error start module 'natsapi':", err)
		}

		d.nats = apiNats
	}

	return d.nats
}

// KafkaConnecter подключение к Kafka
func (d *DiContainer) KafkaConnecter(ctx context.Context) KafkaConnecter {
	if d.kafka == nil {
		apiKafka, err := kafkaapi.New(
			d.Logger(ctx),
			d.Counter(ctx),
			kafkaapi.WithHost(d.Configer().GetKafka().Host),
			kafkaapi.WithPort(d.Configer().GetKafka().Port),
			kafkaapi.WithCacheTTL(d.Configer().GetKafka().CacheTTL),
			kafkaapi.WithTopicsSubscription(d.Configer().GetKafka().Topics),
			kafkaapi.WithNameRegionalObject(d.Configer().GetCommon().RegionalObject),
		)
		if err != nil {
			log.Fatal("error initialization module 'kafkaapi':", err)
		}

		if err = apiKafka.Start(ctx); err != nil {
			log.Fatal("error start module 'kafkaapi':", err)
		}

		d.kafka = apiKafka
	}

	return d.kafka
}
