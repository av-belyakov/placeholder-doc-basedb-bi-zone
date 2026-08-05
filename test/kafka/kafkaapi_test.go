package kafka_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	kafkaapi "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supporting"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func TestKafkaApi(t *testing.T) {
	unsetAllEnviromentEnvAny()

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := supporting.NewLogging()
	counting := &supporting.Counting{}

	go func(ctx context.Context, l *supporting.Logging) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	//инициализация файла для записи входящих сообщений
	f, err := os.OpenFile(
		fmt.Sprintf("incoming_messages_%s.log", time.Now().Format(time.RFC3339)),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//инициализация модуля
	kafkaApiModule, err := kafkaapi.New(
		logging,
		counting,
		kafkaapi.WithHost(cfg.GetKafka().Host),
		kafkaapi.WithPort(cfg.GetKafka().Port),
		kafkaapi.WithCacheTTL(cfg.GetKafka().CacheTTL),
		kafkaapi.WithAuthType(cfg.GetKafka().AuthType),
		kafkaapi.WithSASLMechanism(cfg.GetKafka().SASLMechanism),
		kafkaapi.WithSSLUserName(cfg.GetKafka().SSLUsername),
		kafkaapi.WithSSLPassword(cfg.GetKafka().SSLPassword),
		kafkaapi.WithSSLCeFile(cfg.GetKafka().SSLCaFile),
		kafkaapi.WithTopicsSubscription(cfg.GetKafka().Topics),
		kafkaapi.WithNameRegionalObject(cfg.GetCommon().RegionalObject),
	)
	if err != nil {
		log.Fatal(err)
	}

	//запуск модуля
	if err := kafkaApiModule.StartConsumer(ctx); err != nil {
		log.Fatal(err)
	}

	for msg := range kafkaApiModule.GetChannelFromModule() {
		//t.Logf("Received message:%s\n\n", string(msg.Data))
		str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
		if err != nil {
			t.Logf("Error: %+v\n", err)

			continue
		}

		//t.Logf("Received message:\n%s\n", str)
		if _, err = fmt.Fprintf(f, "--- TOPIC - '%s' %s\n%s\n\n", msg.SubjectType, time.Now().Format(time.RFC3339), str); err != nil {
			t.Logf("Error: %+v\n", err)
		}
	}

	t.Cleanup(func() {
		stop()
		unsetAllEnviromentEnvAny()
	})
}

func unsetAllEnviromentEnvAny() {
	os.Unsetenv("GO_PHDOCBASEDBBZ_MAIN")

	//настройка наименования регионального объекта
	os.Unsetenv("GO_PHDOCBASEDBBZ_REGIONALOBJECT")

	//настройки NATS
	os.Unsetenv("GO_PHDOCBASEDBBZ_NHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NCACHETTL")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NSUBLISTENER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NENRICHINGQUER")

	//настройки Kafka
	os.Unsetenv("GO_PHDOCBASEDBBZ_KHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KTOPICS")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KCACHETTL")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KLOGIN")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KCERTPATH")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KTRUSTSTOREPATH")

	// Настройки доступа к БД в которую будут записыватся alert и case
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGENAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGETABLES")

	//настройки доступа к БД в которую будут записыватся логи
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGNAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME")
}
