package datamodels_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodeltest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneTag(t *testing.T) {
	tag := &datamodels.BiZoneIRPTag{}

	listTesting := map[string]datamodeltest.TestOptions{}

	listTesting["Name"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Name(),
		SetFunc: func() {
			tag.SetAnyName(listTesting["Name"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, tag.GetName(), listTesting["Name"].ValueString)
		},
	}
	listTesting["Color"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Color(),
		SetFunc: func() {
			tag.SetAnyColor(listTesting["Color"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, tag.GetColor(), listTesting["Color"].ValueString)
		},
	}

	listTesting["Created"] = datamodeltest.TestOptions{
		ValueString: gofakeit.ConnectiveTime(),
		SetFunc: func() {
			tag.SetAnyCreated(listTesting["Created"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, tag.GetCreated(), listTesting["Created"].ValueString)
		},
	}

	listTesting["CreatedBy.Username"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Username(),
		SetFunc: func() {
			tag.SetAnyCreatedByUsername(listTesting["CreatedBy.Username"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, tag.GetCreatedByUsername(), listTesting["CreatedBy.Username"].ValueString)
		},
	}

	listTesting["CreatedBy.Id"] = datamodeltest.TestOptions{
		ValueUint64: gofakeit.Uint64(),
		SetFunc: func() {
			tag.SetAnyCreatedByID(listTesting["CreatedBy.Id"].ValueUint64)
		},
		GetFunc: func() {
			assert.Equal(t, tag.GetCreatedByID(), listTesting["CreatedBy.Id"].ValueUint64)
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
