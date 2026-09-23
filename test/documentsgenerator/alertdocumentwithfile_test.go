package documentsgenerator

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
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/countermessage"
)

type ChMessage struct {
	Type    string
	Message string
}

func NewChMessage() *ChMessage {
	return &ChMessage{}
}

func (chm *ChMessage) GetType() string {
	return chm.Type
}

func (chm *ChMessage) GetMessage() string {
	return chm.Message
}

func (chm *ChMessage) SetType(v string) {
	chm.Type = v
}

func (chm *ChMessage) SetMessage(v string) {
	chm.Message = v
}

type Logging struct {
	ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{ch: make(chan interfaces.Messager)}
}

func (l *Logging) SetChan(v chan interfaces.Messager) {
	l.ch = v
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.ch
}

func (l *Logging) Send(msgType, msgData string) {
	msg := NewChMessage()
	msg.SetType(msgType)
	msg.SetMessage(msgData)

	l.ch <- msg
}

func TestAlertDocumentWithFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	alertRaw, err := os.ReadFile("../../test/test_json/alertmgr.json")
	if err != nil {
		log.Fatalln(err)
	}

	logging := NewLogging()
	counting := countermessage.New(logging.ch)
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

	id, verifedBiZoneAlert, listRawFields, err := documentgenerator.BiZoneAlertsGenerator(decoder.Start(alertRaw))
	assert.NoError(t, err)

	fmt.Println("ID:", id)
	fmt.Println("List Raw Fields:")
	for k, v := range listRawFields {
		fmt.Printf("\t%s:%s\n", k, v)
	}

	fmt.Println("\nVerifedBiZoneAlert")
	godump.DumpJSON(verifedBiZoneAlert)

	t.Cleanup(func() {
		cancel()
	})
}
