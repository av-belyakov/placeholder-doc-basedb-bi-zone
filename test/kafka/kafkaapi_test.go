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
	"golang.org/x/sync/errgroup"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/kafkaapi"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)

	logging := supporting.NewLogging()
	counting := &supporting.Counting{}

	doneLogging := make(chan struct{})
	go func(ctx context.Context, l *supporting.Logging) {
		for {
			select {
			case <-doneLogging:
				log.Println("placeholder_doc-basedb-bi-zone login module is Stop")

				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	var fAlert, fCase *os.File

	errGr := errgroup.Group{}
	errGr.Go(func() error {
		//инициализация файла для записи входящих сообщений типа alert
		var err error
		fAlert, err = os.OpenFile(
			fmt.Sprintf("incoming_messages_type_alert_%s.log", time.Now().Format(time.RFC3339)),
			os.O_RDWR|os.O_CREATE,
			0644,
		)

		return err
	})
	errGr.Go(func() error {
		//инициализация файла для записи входящих сообщений типа case
		var err error
		fCase, err = os.OpenFile(
			fmt.Sprintf("incoming_messages_type_case_%s.log", time.Now().Format(time.RFC3339)),
			os.O_RDWR|os.O_CREATE,
			0644,
		)

		return err
	})
	if err := errGr.Wait(); err != nil {
		log.Fatal(err)
	}

	//инициализация модуля
	kafkaApiModule, err := kafkaapi.New(
		logging,
		counting,
		kafkaapi.WithHost(cfg.GetKafka().Host),
		kafkaapi.WithPort(cfg.GetKafka().Port),
		kafkaapi.WithGroupId(cfg.GetKafka().GroupId),
		kafkaapi.WithCacheTTL(cfg.GetKafka().CacheTTL),
		kafkaapi.WithSASLMechanism(cfg.GetKafka().SASLMechanism),
		kafkaapi.WithSSLUserName(cfg.GetKafka().SSLUsername),
		kafkaapi.WithSSLPassword(cfg.GetKafka().SSLPassword),
		kafkaapi.WithSSLCeFile(cfg.GetKafka().SSLCaFile),
		kafkaapi.WithSSLCertFile(cfg.GetKafka().SSLCertFile),
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

	go func() {
		<-ctx.Done()

		if err := kafkaApiModule.StopConsumer(shutdownCtx); err != nil {
			log.Println(fmt.Errorf("kafka consumer shutdown: %w", err).Error())
		}

		close(doneLogging)

		log.Println("placeholder_doc-basedb-bi-zone module is Stop")
	}()

	for msg := range kafkaApiModule.GetChannelFromModule() {

		t.Logf("Received subject type:%s\n", string(msg.SubjectType))

		str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
		if err != nil {
			t.Logf("Error: %+v\n", err)

			continue
		}

		switch msg.SubjectType {
		case "alerts":
			if _, err = fmt.Fprintf(fAlert, "--- TOPIC - '%s' %s\n%s\n\n", msg.SubjectType, time.Now().Format(time.RFC3339), str); err != nil {
				t.Logf("Error: %+v\n", err)
			}

		case "cases":
			if _, err = fmt.Fprintf(fCase, "--- TOPIC - '%s' %s\n%s\n\n", msg.SubjectType, time.Now().Format(time.RFC3339), str); err != nil {
				t.Logf("Error: %+v\n", err)
			}

		default:
			t.Logf("an undefined type message was obtained, type message is '%s'\n", msg.SubjectType)
		}
	}

	t.Cleanup(func() {
		unsetAllEnviromentEnvAny()

		fAlert.Close()
		fCase.Close()

		stop()
		shutdownCancel()
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
