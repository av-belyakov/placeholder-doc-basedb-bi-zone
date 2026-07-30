package datamodels_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestObservedIocs(t *testing.T) {
	observedIocs := &datamodels.BiZoneIRPObservedIocs{}

	listTesting := map[string]datamodelstest.TestOptions{}

	listTesting["Ioc"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			observedIocs.SetAnyIoc(listTesting["Ioc"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, observedIocs.GetIoc(), listTesting["Ioc"].ValueString)
		},
	}

	listTesting["IocType"] = datamodelstest.TestOptions{
		ValueString: gofakeit.AnimalType(),
		SetFunc: func() {
			observedIocs.SetAnyIocType(listTesting["IocType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, observedIocs.GetIocType(), listTesting["IocType"].ValueString)
		},
	}

	listTesting["Category"] = datamodelstest.TestOptions{
		ValueSliceString: gofakeit.NiceColors(),
		SetFunc: func() {
			observedIocs.SetCategories(listTesting["Category"].ValueSliceString)
		},
		GetFunc: func() {
			assert.True(t, slices.Equal(observedIocs.GetCategories(), listTesting["Category"].ValueSliceString))
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
