package databasestorageapi_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/response"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

func TestUpdateSensorInformationToIncident(t *testing.T) {
	const (
		Incident_Key        = "INC-2045"
		Special_UUID        = "dcf51262-0462-4fdc-8691-dff006fa110a" // это для поиска инцидентов
		Elasticsearch_ID    = "4RGIh6EBcfH0hhBlde3x"                 // по этому _id выполняется обновление
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

	go func(ctx context.Context, l interfaces.Logger) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				if msg.GetType() == "error" {
					assert.NoError(t, errors.New(msg.GetMessage()))

					return
				}

				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	apiDBS, err := databasestorageapi.New(
		logging,
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

	t.Run("Добавление в инциденты хранящиеся в БД дополнительной информации о сенсорах", func(t *testing.T) {
		request, err := json.MarshalIndent(response.ResponseSensorsInformation{
			Informations: []response.SensorInformation{
				{
					INN:                      "7703602065",
					GeoCode:                  "RU",
					SensorID:                 "570009",
					ObjectArea:               "Агропромышленный комплекс",
					OrganizationName:         "РусАгроКомплекс",
					FullOrganizationName:     "АО 'РусАгроКомплекс'",
					SubjectRussianFederation: "Производство",
				},
			},
			Source: "test-source",
			TaskId: Elasticsearch_Index + ":" + Special_UUID,
		}, "", " ")
		assert.NoError(t, err)

		err = apiDBS.AddSensorInformation_Test(ctx, request)
		assert.NoError(t, err)

		/*
			sensorInfoList := []datamodels.SensorInformation{
				{
					INN:         "7703602065",
					HostId:      "1234", //SpecialSensorID
					GeoCode:     "RU",
					OrgName:     "РусАгроКомплекс", //OrganizationName:
					SensorId:    "570009",          //SensorID
					SubjectRF:   "Производство",    //SubjectRussianFederation
					ObjectArea:  "Агропромышленный комплекс",
					FullOrgName: "АО 'РусАгроКомплекс'", //FullOrganizationName

				},
			}

			//обновление информации в БД
			request, err := json.MarshalIndent(datamodels.AdditionalInformationSensors{Sensors: sensorInfoList}, "", " ")
			assert.NoError(t, err)

			bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
			client := apiDBS.GetClientDB()

			res, err := client.Update(Elasticsearch_Index, Elasticsearch_ID, bodyUpdate)
			//res, err := client.Update(Elasticsearch_Index, Special_UUID, bodyUpdate)
			assert.NoError(t, err)
			assert.Equal(t, res.StatusCode, http.StatusOK)
			assert.NotNil(t, res)
			assert.NotNil(t, res.Body)

			resInfo, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			fmt.Println("Response:", string(resInfo))

			res.Body.Close()
		*/
	})

	t.Run("Добавление в инциденты хранящиеся в БД дополнительной информации об ip адресах", func(t *testing.T) {
		var geoIpInfo []datamodels.IpAddressInformation = []datamodels.IpAddressInformation{
			{
				Ip:          "217.77.55.94",
				City:        "Havana",
				Country:     "Kuba",
				CountryCode: "CU",
			},
		}

		request, err := json.MarshalIndent(datamodels.AdditionalInformationIpAddress{IpAddresses: geoIpInfo}, "", " ")
		assert.NoError(t, err)

		err = apiDBS.AddGeoIPInformation_Test(ctx, request)
		assert.NoError(t, err)

		// ищем объект с таким же идентификатором как и принятый в обработку объект
		res, err := apiDBS.GetDocument(
			ctx,
			[]string{Elasticsearch_Index},
			strings.NewReader(
				fmt.Sprintf(
					`{
					"query": {
						"bool": {
							"must": [{"match": {"key": "%s"}}]}}}`,
					Incident_Key)),
			0)
		assert.NoError(t, err)

		response := databasestorageapi.ResponseVerifiedIRPIncidents{}
		err = json.Unmarshal(res, &response)
		assert.NoError(t, err)
		assert.Greater(t, response.Options.Total.Value, 0)

		fmt.Println("Search incident responce:")
		for k, v := range response.Options.Hits {
			fmt.Printf("%d.\n\t%+v\n", k, v.Source)
		}

		/*
			//выполняется обновление информации в БД
			bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))

			client := apiDBS.GetClientDB()
			res, err := client.Update(Elasticsearch_Index, Elasticsearch_ID, bodyUpdate)
			assert.NoError(t, err)
			assert.Equal(t, res.StatusCode, http.StatusOK)
			assert.NotNil(t, res)
			assert.NotNil(t, res.Body)

			resInfo, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			fmt.Println("Response:", string(resInfo))

			res.Body.Close()
		*/
	})
}

/*
	{
		inn:         "7703602065",
		host_id:      "1234",
		geo_code:     "RU",
		org_name:     "РусАгроКомплекс",
		sensor_id:    "570009",
		subject_rf:   "Производство",
		object_area:  "Агропромышленный комплекс",
		full_org_name: "АО 'РусАгроКомплекс'",
	}
*/
