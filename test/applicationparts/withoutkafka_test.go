package applicationparts

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/irpapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supporting"
	"github.com/joho/godotenv"
)

func TestAppWithOutKafka(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	apiIrp, err := irpapi.New(
		irpapi.WithAPIKey(cfg.GetIRP().ApiToken),
		irpapi.WithHost(cfg.GetIRP().Host),
		irpapi.WithPort(cfg.GetIRP().Port),
		irpapi.WithConnectionTimeout(cfg.GetIRP().ConnectionTimeout),
	)
	if err != nil {
		log.Fatal("initialization module 'irpapi':", err)
	}

	logger := supporting.NewLogging()
	apiDBS, err := databasestorageapi.New(
		logger,
		databasestorageapi.WithHost(cfg.GetStorageDB().Host),
		databasestorageapi.WithPort(cfg.GetStorageDB().Port),
		databasestorageapi.WithNameDB(cfg.GetStorageDB().NameDB),
		databasestorageapi.WithUser(cfg.GetStorageDB().User),
		databasestorageapi.WithPasswd(cfg.GetStorageDB().Passwd),
		databasestorageapi.WithStorage(cfg.GetStorageDB().Storage),
		databasestorageapi.WithMaxGetDocumentSize(15))
	if err != nil {
		log.Fatal("initialization module 'db':", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err = apiDBS.Start(ctx)
	if err != nil {
		log.Fatal("start module 'db':", err)
	}

	chDone := make(chan struct{})
	go func() {
		for {
			select {
			case <-chDone:
				log.Println("placeholder_doc-basedb-bi-zone login module is Stop")

				return

			case msg := <-logger.GetChan():
				fmt.Printf("LOG: %s %s\n", msg.GetType(), msg.GetMessage())

			case <-apiDBS.GetChannelFromModule():
				// так как нужно принимать из модуля запросы на обогащение информации

			}
		}
	}()

	ch := make(chan string)
	err = getIncidentsId(ctx, "./incidentsid.txt", 3, ch)
	if err != nil {
		log.Fatal(err)
	}

	var (
		wg                     sync.WaitGroup
		numberIncidentsHandled int
	)

	wg.Go(func() {
		decoder := decoderjsondocuments.New(logger)

		for {
			select {
			case <-ctx.Done():
				return

			case id, ok := <-ch:
				if !ok {
					return
				}

				time.Sleep(time.Second)

				if ctx.Err() != nil {
					return
				}

				fmt.Printf("получаем полную информацию по инциденту '%s'\n", id)
				response, statusCode, err := apiIrp.GetIncident(ctx, id)
				if err != nil {
					fmt.Println(err)

					continue
				}

				if statusCode != http.StatusOK {
					fmt.Printf("принят код статуса %d %s\n", statusCode, string(response))

					continue
				}

				fmt.Printf("генерируем проверенную структуру для инцидента '%s'\n", id)
				_, verifiedBiZoneIncident, _, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(response))
				if err != nil {
					fmt.Println(err)

					continue
				}

				fmt.Printf("загрузка инцидента '%s' в базу данных\n", id)
				err = apiDBS.TestAddIRPIncident(ctx, verifiedBiZoneIncident)
				if err != nil {
					fmt.Println(err)
				}

				numberIncidentsHandled++
			}
		}
	})

	wg.Wait()

	chDone <- struct{}{}

	fmt.Println("Number incidents handled:", numberIncidentsHandled)

	t.Cleanup(func() {
		unsetAllEnviromentEnvAny()
		stop()
	})
}

// getIncidentsId читает файл со списком инцидентов и по очереди передает id каждого инцидента
func getIncidentsId(ctx context.Context, file string, count int, ch chan<- string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	go func() {
		defer func() {
			f.Close()
			close(ch)
		}()

		var num int
		for {
			select {
			case <-ctx.Done():
				return

			default:
				if scanner.Scan() {
					if count != 0 && num == count {
						return
					}

					num++
					ch <- scanner.Text()
				}

				// Проверяем, не было ли ошибки при сканировании
				if err := scanner.Err(); err != nil {
					fmt.Printf("чтение файла %s: %s", file, err.Error())
				}
			}
		}
	}()

	return nil
}

func unsetAllEnviromentEnvAny() {
	os.Unsetenv("GO_PHDOCBASEDBBZ_MAIN")

	os.Unsetenv("GO_PHDOCBASEDBBZ_IRPAPITOKEN")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD")
}
