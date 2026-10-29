package databasestorageapi_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestSearchUnderlineId(t *testing.T) {
	const (
		Special_UUID        = "dcf51262-0462-4fdc-8691-dff006fa110a"
		Elasticsearch_Index = "module_placeholderdb_bizone_incidents.test_2026_10"
	)

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("GO_PHDOCBASEDBBZ_DBWLOGPASSWD =", os.Getenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD"))
	fmt.Println("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD =", os.Getenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := helpers.NewLogging()
	counting := helpers.NewCounting()

	go func(ctx context.Context, l interfaces.Logger, c *helpers.Counting) {
		for {
			select {
			case <-ctx.Done():
				return

			case count := <-c.GetChan():
				fmt.Println("count.message:", count.Message, " count.Count", count.Count)

			case msg := <-l.GetChan():
				if msg.GetType() == "error" {
					assert.NoError(t, errors.New(msg.GetMessage()))

					return
				}

				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging, counting)

	apiDBS, err := databasestorageapi.New(
		logging,
		counting,
		databasestorageapi.WithHost(cfg.StorageDB.Host),
		databasestorageapi.WithPort(cfg.StorageDB.Port),
		databasestorageapi.WithNameDB(cfg.StorageDB.NameDB),
		databasestorageapi.WithUser(cfg.StorageDB.User),
		databasestorageapi.WithPasswd(cfg.StorageDB.Passwd),
		databasestorageapi.WithStorage(cfg.StorageDB.Storage),
		databasestorageapi.WithMaxGetDocumentSize(15))
	if err != nil {
		log.Fatalln(err)
	}

	if err := apiDBS.Start(ctx); err != nil {
		log.Fatalln(err)
	}

	t.Run("Поиск инцидента по его @special_uuid", func(t *testing.T) {
		underlineId, err := apiDBS.GetUnderlineId(ctx, Elasticsearch_Index, Special_UUID)
		assert.NoError(t, err)
		assert.NotEqual(t, underlineId, "")
	})

	/*t.Run("Поиск информации GeoIP", func(t *testing.T) {
		newGeoIpList := []datamodels.IpAddressInformation{
			{Ip: "45.13.191.34", City: "Oslolo", Country: "Норвегия", CountryCode: "NO"},
			{Ip: "45.13.191.28", City: "Oslolo", Country: "Норвегия", CountryCode: "NO"},
			{Ip: "45.13.191.18", City: "Oslo", Country: "Норвегия", CountryCode: "NO"},
			{Ip: "45.13.191.47", City: "Oslo", Country: "Норвегия", CountryCode: "NO"},
			{Ip: "98.113.101.17", City: "Xmu", Country: "Нигерия", CountryCode: "NG"},
		}

		underlineId, geoIpList, err := apiDBS.SearchGeoIPInformation(ctx, Elasticsearch_Index, Special_UUID)
		assert.NoError(t, err)
		assert.Equal(t, underlineId, "4g0WIJcBRHX25kGeUOOR")

		t.Log("underlineId:", underlineId)
		fmt.Println("BEFORE @ipAddressAdditionalInformation:")
		for k, v := range geoIpList {
			fmt.Printf("%d. %#v\n", k, v)
		}

		for k, v := range newGeoIpList {
			num, ok := supportingfunctions.SliceContainsElementFunc(geoIpList, func(num int) bool {
				if v.Ip == geoIpList[num].Ip {
					return true
				}

				return false
			})
			if ok {
				geoIpList[num] = v
			} else {
				geoIpList = append(geoIpList, v)
			}

			t.Logf("ok:%t, k:%d, num:%d\n", ok, k, num)
		}

		fmt.Println("AFTER @ipAddressAdditionalInformation:")
		for k, v := range geoIpList {
			fmt.Printf("%d. %#v\n", k, v)
		}
	})*/

	t.Cleanup(func() {
		stop()

		os.Unsetenv("GO_PHDOCBASEDBBZ_MAIN")
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")
	})
}
