package datamodels_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodeltest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneWatcher(t *testing.T) {
	watcher := &datamodels.BiZoneIRPWatcher{}

	listTesting := map[string]datamodeltest.TestOptions{}

	listTesting["ID"] = datamodeltest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			watcher.SetAnyId(listTesting["ID"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetId(), listTesting["ID"].ValueString)
		},
	}

	listTesting["Username"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Name(),
		SetFunc: func() {
			watcher.SetAnyUsername(listTesting["Username"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetUsername(), listTesting["Username"].ValueString)
		},
	}

	listTesting["FirstName"] = datamodeltest.TestOptions{
		ValueString: gofakeit.FirstName(),
		SetFunc: func() {
			watcher.SetAnyFirstName(listTesting["FirstName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetFirstName(), listTesting["FirstName"].ValueString)
		},
	}

	listTesting["LastName"] = datamodeltest.TestOptions{
		ValueString: gofakeit.LastName(),
		SetFunc: func() {
			watcher.SetAnyLastName(listTesting["LastName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetLastName(), listTesting["LastName"].ValueString)
		},
	}

	listTesting["Email"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Email(),
		SetFunc: func() {
			watcher.SetAnyEmail(listTesting["Email"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetEmail(), listTesting["Email"].ValueString)
		},
	}

	listTesting["Patronimic"] = datamodeltest.TestOptions{
		ValueString: gofakeit.HipsterParagraph(),
		SetFunc: func() {
			watcher.SetAnyPatronimic(listTesting["Patronimic"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetPatronimic(), listTesting["Patronimic"].ValueString)
		},
	}

	listTesting["IsActive"] = datamodeltest.TestOptions{
		ValueBool: gofakeit.Bool(),
		SetFunc: func() {
			watcher.SetAnyIsActive(listTesting["IsActive"].ValueBool)
		},
		GetFunc: func() {
			assert.Equal(t, watcher.GetIsActive(), listTesting["IsActive"].ValueBool)
		},
	}

	var num int
	for k, v := range listTesting {
		num++
		t.Run(fmt.Sprintf("Test %d. Field %s", num, k), func(t *testing.T) {
			v.SetFunc()
			v.GetFunc()
		})
	}
}
