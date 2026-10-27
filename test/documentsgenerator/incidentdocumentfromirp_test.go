package documentsgenerator_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/goforj/godump"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/irpapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/documentsgenerator"
)

func TestGetIncidentDocumentFromIRP(t *testing.T) {
	const (
		Incident_Number string = "INC-2039"
	)

	var (
		data       []byte
		statusCode int
		err        error
	)

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(t.Context())

	logging := documentsgenerator.NewLogging()
	counting := countermessage.New(logging.GetChanToData())
	counting.Start(ctx)

	decoder := decoderjsondocuments.New(counting, logging)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-logging.GetChan():
				fmt.Println("Log:", msg)

			}
		}
	}()

	apiToken := os.Getenv("GO_PHDOCBASEDBBZ_IRPAPITOKEN")
	fmt.Println("Api Token:", apiToken)

	api, err := irpapi.New(
		irpapi.WithAPIKey(apiToken),
		irpapi.WithHost(cfg.IRP.Host),
		irpapi.WithPort(cfg.IRP.Port),
		irpapi.WithConnectionTimeout(7),
	)
	if err != nil {
		log.Fatalln(err)
	}

	t.Run(fmt.Sprintf("Запрашиваем у IRP данные по инциденту №%s", Incident_Number), func(t *testing.T) {
		data, statusCode, err = api.GetIncident(t.Context(), Incident_Number)
		assert.NoError(t, err)
		assert.Equal(t, statusCode, http.StatusOK)
	})

	t.Run(fmt.Sprintf("Обрабатываем полученные от IRP данные по инциденту №%s", Incident_Number), func(t *testing.T) {
		uuid, verifiedIncedent, listRawFields, err := documentgenerator.GeneratingVerifiedDocumentFromIncident(decoder.Start(data))
		assert.NoError(t, err)

		fmt.Printf("Document generated UUID:%s\n", uuid)
		fmt.Println("listRawFields:")
		for k, v := range listRawFields {
			fmt.Printf("%s %s\n", k, v)
		}
		fmt.Println("Verified Document:")
		godump.DumpJSON(verifiedIncedent)

	})

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDBBZ_IRPAPITOKEN")
		cancel()
	})
}
