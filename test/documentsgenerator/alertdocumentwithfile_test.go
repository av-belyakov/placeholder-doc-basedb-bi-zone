package documentsgenerator_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/goforj/godump"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/documentsgenerator"
)

func TestAlertDocumentWithFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	alertRaw, err := os.ReadFile("../../test/test_json/alertmgr.json")
	if err != nil {
		log.Fatalln(err)
	}

	logging := documentsgenerator.NewLogging()

	decoder := decoderjsondocuments.New(logging)
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

	id, verifedBiZoneAlert, listRawFields, err := documentgenerator.BiZoneAlertsGenerator(decoder.Start(alertRaw))
	assert.NoError(t, err)

	fmt.Println("\nID:", id)
	fmt.Println("List Raw Fields:")
	for k, v := range listRawFields {
		fmt.Printf("\t%s:%s\n", k, v)
	}

	fmt.Println("VerifedBiZoneAlert")
	godump.DumpJSON(verifedBiZoneAlert)

	t.Cleanup(func() {
		cancel()
	})
}
