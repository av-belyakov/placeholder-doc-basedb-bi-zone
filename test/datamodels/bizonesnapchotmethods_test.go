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

func TestBiZoneSnapShot(t *testing.T) {
	snapshot := &datamodels.BiZoneIRPSnapshot{}

	listTesting := map[string]datamodeltest.TestOptions{}

	listTesting["OS"] = datamodeltest.TestOptions{
		ValueString: gofakeit.OperaUserAgent(),
		SetFunc: func() {
			snapshot.SetAnyOS(listTesting["OS"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetOS(), listTesting["OS"].ValueString)
		},
	}
	listTesting["Fqdn"] = datamodeltest.TestOptions{
		ValueString: gofakeit.URL(),
		SetFunc: func() {
			snapshot.SetAnyFqdn(listTesting["Fqdn"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetFqdn(), listTesting["Fqdn"].ValueString)
		},
	}

	listTesting["Domain"] = datamodeltest.TestOptions{
		ValueString: gofakeit.DomainName(),
		SetFunc: func() {
			snapshot.SetAnyDomain(listTesting["Domain"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetDomain(), listTesting["Domain"].ValueString)
		},
	}

	listTesting["CMDBID"] = datamodeltest.TestOptions{
		ValueString: gofakeit.CarModel(),
		SetFunc: func() {
			snapshot.SetAnyCMDBID(listTesting["CMDBID"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetCMDBID(), listTesting["CMDBID"].ValueString)
		},
	}

	listTesting["OSType"] = datamodeltest.TestOptions{
		ValueString: gofakeit.BankType(),
		SetFunc: func() {
			snapshot.SetAnyOSType(listTesting["OSType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetOSType(), listTesting["OSType"].ValueString)
		},
	}

	listTesting["Hostname"] = datamodeltest.TestOptions{
		ValueString: gofakeit.DomainName(),
		SetFunc: func() {
			snapshot.SetAnyHostname(listTesting["Hostname"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetHostname(), listTesting["Hostname"].ValueString)
		},
	}

	listTesting["UserCMDBName"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Username(),
		SetFunc: func() {
			snapshot.SetAnyUserCMDBName(listTesting["UserCMDBName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetUserCMDBName(), listTesting["UserCMDBName"].ValueString)
		},
	}

	listTesting["IPAddresses"] = datamodeltest.TestOptions{
		ValueSliceString: []string{
			gofakeit.IPv4Address(),
			gofakeit.IPv4Address(),
			gofakeit.IPv4Address(),
		},
		SetFunc: func() {
			for _, v := range listTesting["IPAddresses"].ValueSliceString {
				snapshot.SetAnyIPAddresse(v)
			}
		},
		GetFunc: func() {
			assert.True(t, slices.Equal(snapshot.GetIPAddresses(), listTesting["IPAddresses"].ValueSliceString))

			snapshot.SetAnyIPAddresse(listTesting["IPAddresses"].ValueSliceString[0])
			snapshot.SetAnyIPAddresse(listTesting["IPAddresses"].ValueSliceString[1])
			assert.True(t, slices.Equal(snapshot.GetIPAddresses(), listTesting["IPAddresses"].ValueSliceString))
		},
	}

	listTesting["MACAddresses"] = datamodeltest.TestOptions{
		ValueSliceString: []string{
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
			gofakeit.MacAddress(),
		},
		SetFunc: func() {
			for _, v := range listTesting["MACAddresses"].ValueSliceString {
				snapshot.SetAnyMACAddresse(v)
			}
		},
		GetFunc: func() {
			assert.True(t, slices.Equal(snapshot.GetMACAddresses(), listTesting["MACAddresses"].ValueSliceString))

			snapshot.SetAnyMACAddresse(listTesting["MACAddresses"].ValueSliceString[0])
			snapshot.SetAnyMACAddresse(listTesting["MACAddresses"].ValueSliceString[1])
			assert.True(t, slices.Equal(snapshot.GetMACAddresses(), listTesting["MACAddresses"].ValueSliceString))
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
