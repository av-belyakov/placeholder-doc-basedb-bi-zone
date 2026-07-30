package datamodels_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
)

func TestBiZoneTypeRefMethods(t *testing.T) {
	typeRef := &datamodels.BiZoneIRPTypeRef{}

	listTesting := map[string]struct {
		Value   string
		SetFunc func()
		GetFunc func()
	}{}

	listTesting["ID"] = struct {
		Value   string
		SetFunc func()
		GetFunc func()
	}{
		Value: gofakeit.ID(),
		SetFunc: func() {
			typeRef.SetAnyId(listTesting["ID"].Value)
		},
		GetFunc: func() {
			assert.Equal(t, typeRef.GetId(), listTesting["ID"].Value)
		},
	}

	listTesting["Title"] = struct {
		Value   string
		SetFunc func()
		GetFunc func()
	}{
		Value: gofakeit.BookTitle(),
		SetFunc: func() {
			typeRef.SetAnyTitle(listTesting["Title"].Value)
		},
		GetFunc: func() {
			assert.Equal(t, typeRef.GetTitle(), listTesting["Title"].Value)
		},
	}

	listTesting["Description"] = struct {
		Value   string
		SetFunc func()
		GetFunc func()
	}{
		Value: gofakeit.ProductDescription(),
		SetFunc: func() {
			typeRef.SetAnyDescription(listTesting["Description"].Value)
		},
		GetFunc: func() {
			assert.Equal(t, typeRef.GetDescription(), listTesting["Description"].Value)
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
