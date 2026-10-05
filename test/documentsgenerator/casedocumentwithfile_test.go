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
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/documentsgenerator"
)

func TestCaseDocumentWithFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	caseRaw, err := os.ReadFile("../../test/test_json/bizonecase.json")
	if err != nil {
		log.Fatalln(err)
	}

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

	id, verifedBiZoneCase, listRawFields, err := documentgenerator.BiZoneCasesGenerator(decoder.Start(caseRaw))
	assert.NoError(t, err)

	fmt.Println("\nID:", id)
	fmt.Println("List Raw Fields:")
	for k, v := range listRawFields {
		fmt.Printf("\t%s:%s\n", k, v)
	}

	fmt.Println("VerifedBiZoneCase")
	godump.DumpJSON(verifedBiZoneCase)

	t.Cleanup(func() {
		cancel()
	})
}
