package irpapi

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/irpapi"
)

func TestIRPAPIRequests(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	cfg, err := confighandler.New(constants.Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

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

	data, statusCode, err := api.GetIncident(t.Context(), "INC-2040")
	assert.NoError(t, err)
	assert.Equal(t, statusCode, http.StatusOK)

	fmt.Printf("DATA:'%s'\n", string(data))

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDBBZ_IRPAPITOKEN")
	})
}
