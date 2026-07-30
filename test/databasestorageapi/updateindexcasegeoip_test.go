package databasestorageapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/helpers"
)

const (
	Index_Test         = "test.module_placeholder_case_index"
	Filepath_Test_Case = "../test_json/case_39100.json"
)

var (
	listIPAddress []datamodels.IpAddressInformation = []datamodels.IpAddressInformation{
		{
			Ip:          "96.136.64.9",
			City:        "Havana",
			Country:     "Kuba",
			CountryCode: "CU",
		},
		{
			Ip:          "72.31.66.61",
			City:        "Sidney",
			Country:     "Australia",
			CountryCode: "AU",
		},
		{
			Ip:          "13.22.63.6",
			City:        "Lida",
			Country:     "Livia",
			CountryCode: "LI",
		},
	}

	logging  *helpers.Logging
	counting *helpers.Counting
)

func CreateTestCase(ctx context.Context, filePath string) (string, *datamodels.VerifiedBiZoneIRPAlert, error) {
	var (
		rootId        string
		verifiedAlert *datamodels.VerifiedBiZoneIRPAlert
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	b, err := os.ReadFile(filePath)
	if err != nil {
		return rootId, verifiedAlert, err
	}

	decoder := decoderjsondocuments.New(counting, logging)
	chDecode := decoder.Start(b)

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
					log.Fatal(msg.GetMessage())

					return
				}

				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging, counting)

	rootId, verifiedAlert, _ = documentgenerator.BiZoneAlertsGenerator(chDecode)

	return rootId, verifiedAlert, nil
}

func TestUpdateIndexCaseGeoIp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal(err)
	}

	//подключение к БД
	apiDBS, err := helpers.DataBaseConnectionInitialization(
		ctx,
		logging,
		counting,
		databasestorageapi.WithHost("datahook.cloud.gcm"),
		databasestorageapi.WithPort(9200),
		databasestorageapi.WithUser("writer"),
		databasestorageapi.WithPasswd(os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")),
		databasestorageapi.WithStorage(map[string]string{
			"alert": "module_placeholderdb_alert",
			"case":  "module_placeholderdb_case",
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	//удаление тестового индекса
	//if err := apiDBS.DelIndexSetting(ctx, []string{Index_Test}); err != nil {
	//	t.Log("ERROR:", err)
	//}

	indexes, err := apiDBS.GetExistingIndexes(ctx, Index_Test)
	assert.NoError(t, err)
	assert.Equal(t, len(indexes), 0)

	// получаем тестовый кейс
	rootId, verifyCase, err := CreateTestCase(ctx, Filepath_Test_Case)
	if err != nil {
		t.Fatal(err)
	}

	vcb, err := json.Marshal(verifyCase)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(vcb))

	//добавляем новый тестовый кейс
	statusCode, err := apiDBS.InsertDocument(ctx, Index_Test, vcb)
	assert.NoError(t, err)
	assert.Equal(t, statusCode, 201)

	t.Log("Insert document, status code:", statusCode)

	//устанавливаем максимальный лимит полей в 2000
	err = apiDBS.SetMaxTotalFieldsLimit(ctx, []string{Index_Test})
	assert.NoError(t, err)

	//проверяем наличие кейса
	underlineId, err := apiDBS.GetUnderlineId(ctx, Index_Test, rootId)
	assert.NoError(t, err)

	t.Log("underlineId:", underlineId)

	request, err := json.MarshalIndent(datamodels.AdditionalInformationIpAddress{
		IpAddresses: listIPAddress,
	}, "", " ")
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	data, statusCode, err := apiDBS.Update(Index_Test, underlineId, bodyUpdate)
	assert.NoError(t, err)
	assert.Equal(t, statusCode, 200)

	fmt.Println("DATA:", data)

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}

/*
	t.Run("", func(t *testing.T) {

	})
*/
