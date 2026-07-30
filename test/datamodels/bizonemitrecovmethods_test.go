package datamodels_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodeltest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneMitreCoveMethods(t *testing.T) {
	mitreCov := &datamodels.BiZoneIRPMITRECOV{}

	listTesting := map[string]datamodeltest.TestOptions{}

	listTesting["Persistence"] = datamodeltest.TestOptions{
		ValueSliceString: gofakeit.NiceColors(),
		SetFunc: func() {
			mitreCov.SetPersistencies(listTesting["Persistence"].ValueSliceString)
		},
		GetFunc: func() {
			assert.True(t, slices.Equal(mitreCov.GetPersistence(), listTesting["Persistence"].ValueSliceString))
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
