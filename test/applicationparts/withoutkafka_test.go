package applicationparts

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/irpapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supporting"
)

const Count_Incident_Id = 10

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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	logger := supporting.NewLogging()

	// инициализация модуля подключения к БД
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
	err = apiDBS.Start(ctx)
	if err != nil {
		log.Fatal("start module 'db':", err)
	}

	// инициализация модуля подключения к NATS
	apiNats, err := natsapi.New(
		logger,
		natsapi.WithHost(cfg.GetNATS().Host),
		natsapi.WithPort(cfg.GetNATS().Port),
		natsapi.WithCacheTTL(cfg.GetNATS().CacheTTL),
		natsapi.WithRequests(cfg.GetNATS().EnrichingQueries),
		natsapi.WithSubscriptions(cfg.GetNATS().Subscriptions),
		natsapi.WithNameRegionalObject(cfg.GetCommon().RegionalObject),
	)
	if err != nil {
		log.Fatal("initialization module 'natsapi' ", err)
	}

	var wg sync.WaitGroup
	wg.Add(Count_Incident_Id * 2)

	chDoneRoute := make(chan struct{})
	chDoneLog := make(chan struct{})
	go func() {
		for {
			select {
			case <-chDoneLog:
				log.Println("\tfile test - placeholder_doc-basedb-bi-zone login module is Stop")

				return

			case msg := <-logger.GetChan():
				fmt.Printf("LOG: %s - %s\n", strings.ToUpper(msg.GetType()), msg.GetMessage())

			}
		}

	}()
	go func() {
		fmt.Println("==== START гороутина обработки данных приходящих через каналы")
		defer fmt.Println("==== END гороутина обработки данных приходящих через каналы")

		for {
			select {
			case <-chDoneRoute:
				log.Println("\tfile test - placeholder_doc-basedb-bi-zone login module is Stop")

				return

			case msg := <-apiDBS.GetChannelFromModule():
				// так как нужно принимать из модуля запросы на обогащение информации
				apiNats.GetChannelToModule() <- natsapi.SettingsChanInput{
					Command: msg.Command,
					RootId:  msg.Id,
					Data:    msg.Data,
				}

				fmt.Printf("\tfile test - получено сообщение от БД, command:'%s', id:'%s', data:'%s'\n", msg.Command, msg.Id, string(msg.Data))

			case msg := <-apiNats.GetChannelFromModule():
				switch msg.SubjectType {
				case "geoip information":
					// передача информации о географическом местоположении ip адресов
					if err := apiDBS.AddGeoIPInformation_Test(ctx, msg.Data); err != nil {
						fmt.Println("geoip information", err.Error())
					}

				case "sensor information":
					// местоположении и принадлежности сенсоров
					if err := apiDBS.AddSensorInformation_Test(ctx, msg.Data); err != nil {
						fmt.Println("sensor information", err.Error())
					}

				default:
					fmt.Println("\tfile test - Undefined subscription type")
				}

				wg.Done()
			}
		}
	}()

	if err = apiNats.Start(ctx); err != nil {
		log.Fatal("start module 'natsapi' ", err)
	}

	ch := make(chan string)
	var numberIncidentsHandled int
	//wg.Go(func() {})
	go func() {
		//fmt.Println("гороутина запроса ID инцидентов")

		numberIncidentsHandled = handlerIncidentId(ctx, ch, logger, apiIrp, apiDBS)
	}()

	err = getIncidentsId(ctx, "./incidentsid.txt", Count_Incident_Id, ch)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	time.Sleep(time.Second * 3)

	chDoneLog <- struct{}{}
	chDoneRoute <- struct{}{}

	fmt.Printf("\tfile test - Number incidents handled: %d\n", numberIncidentsHandled)

	/****************************************

		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	для инцидентов успешно обновляется поле GeoIP, информация о сенсарах не обновляется
	возможно нет инфы по сенсорам, надо проверить

	в обработчике алертов так же ошибка с формированием id запроса @special_uuid что и в инцидентах
	и надо убрать defer и использовать гороутину
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	*****************************************/

	t.Cleanup(func() {
		unsetAllEnviromentEnvAny()
		stop()
	})
}

// handlerIncidentId обработчик инцидентов по их ID
func handlerIncidentId(
	ctx context.Context,
	ch chan string,
	logger *supporting.Logging,
	apiIrp *irpapi.IrpApi,
	apiDBS *databasestorageapi.DatabaseStorage,
) int {
	var count int
	decoder := decoderjsondocuments.New(logger)

	for {
		select {
		case <-ctx.Done():
			//fmt.Println("+++ context done count incident:", count)

			return count

		case id, ok := <-ch:
			//fmt.Printf("=== status chan:'%t'\n", ok)

			if !ok {
				//fmt.Println("+++ close channel count incident:", count)

				return count
			}

			time.Sleep(time.Second)

			if ctx.Err() != nil {
				//fmt.Println("+++ ctx.Err count incident:", count)

				return count
			}

			//fmt.Printf("получаем полную информацию по инциденту '%s'\n", id)
			response, statusCode, err := apiIrp.GetIncident(ctx, id)
			if err != nil {
				fmt.Println(err)

				continue
			}

			if statusCode != http.StatusOK {
				//fmt.Printf("принят код статуса %d %s\n", statusCode, string(response))

				continue
			}

			//fmt.Printf("генерируем проверенную структуру для инцидента '%s'\n", id)
			_, verifiedBiZoneIncident, _, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(response))
			if err != nil {
				fmt.Println("ERROR", err)

				continue
			}

			//fmt.Printf("загрузка инцидента '%s' в базу данных\n", id)
			err = apiDBS.TestAddIRPIncident(ctx, verifiedBiZoneIncident)
			if err != nil {
				fmt.Println("загрузка инцидента", err)
			}

			count++
		}
	}
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
