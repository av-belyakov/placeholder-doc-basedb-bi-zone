package databasestorageapi_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/documentsgenerator"
)

func TestUpdateIncidentFromFile(t *testing.T) {
	const (
		Incident_File_Name = "incident-key-inc-2105.json"
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	incidentRaw, err := os.ReadFile("../../test/test_json/" + Incident_File_Name)
	if err != nil {
		log.Fatalln(err)
	}

	rootPath, err := supportingfunctions.GetPathRoot(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
	}

	if err = os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test"); err != nil {
		log.Fatal(err)
	}

	cfg, err := confighandler.New(rootPath)
	if err != nil {
		log.Fatal("start module 'confighandler':", err)
	}

	logging := documentsgenerator.NewLogging()

	apiDBS, err := databasestorageapi.New(
		logging,
		databasestorageapi.WithHost(cfg.GetStorageDB().Host),
		databasestorageapi.WithPort(cfg.GetStorageDB().Port),
		databasestorageapi.WithNameDB(cfg.GetStorageDB().NameDB),
		databasestorageapi.WithUser(cfg.GetStorageDB().User),
		databasestorageapi.WithPasswd(cfg.GetStorageDB().Passwd),
		databasestorageapi.WithStorage(cfg.GetStorageDB().Storage),
		databasestorageapi.WithMaxGetDocumentSize(15))
	if err != nil {
		log.Fatal("start module 'databasestorageapi':", err)
	}

	if err = apiDBS.Start(ctx); err != nil {
		log.Fatal("start module 'databasestorageapi':", err)
	}

	decoder := decoderjsondocuments.New(logging)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-logging.GetChan():
				if strings.Contains(msg.GetMessage(), "update field") {
					fmt.Printf("Log:\n\ttype:'%s'\n\tmessage:'%s'\n", msg.GetType(), msg.GetMessage())
				}

			case <-apiDBS.GetChannelFromModule():

			}
		}
	}()

	chData := make(chan any, 1)

	t.Run("Тест 1. Генерирование проверенного документа содержащего данные по инциденту", func(t *testing.T) {
		id, verifedData, listRawFields, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(incidentRaw))
		assert.NoError(t, err)

		fmt.Println("\nID:", id)
		fmt.Println("Size list raw fields:", len(listRawFields))

		chData <- verifedData
	})

	t.Run("Тест 2. Добавление инцидента в БД", func(t *testing.T) {
		apiDBS.TestAddIRPIncident(ctx, <-chData)
	})

	/*
		for k, v := range listRawFields {
			fmt.Printf("\t%s:%s\n", k, v)
		}

		fmt.Println("Verifed incident from IRP")
		godump.DumpJSON(verifedData)
	*/

	t.Cleanup(func() {
		close(chData)
		cancel()
	})

}
