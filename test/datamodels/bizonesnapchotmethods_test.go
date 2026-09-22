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

	// --- os ---
	listTesting["OS"] = datamodeltest.TestOptions{
		ValueString: gofakeit.OperaUserAgent(),
		SetFunc: func() {
			snapshot.SetAnyOS(listTesting["OS"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetOS(), listTesting["OS"].ValueString)
		},
	}

	// --- fqdn ---
	listTesting["Fqdn"] = datamodeltest.TestOptions{
		ValueString: gofakeit.URL(),
		SetFunc: func() {
			snapshot.SetAnyFqdn(listTesting["Fqdn"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetFqdn(), listTesting["Fqdn"].ValueString)
		},
	}

	// --- title ---
	listTesting["Title"] = datamodeltest.TestOptions{
		ValueString: gofakeit.BookTitle(),
		SetFunc: func() {
			snapshot.SetAnyTitle(listTesting["Title"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetTitle(), listTesting["Title"].ValueString)
		},
	}

	// --- domain ---
	listTesting["Domain"] = datamodeltest.TestOptions{
		ValueString: gofakeit.DomainName(),
		SetFunc: func() {
			snapshot.SetAnyDomain(listTesting["Domain"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetDomain(), listTesting["Domain"].ValueString)
		},
	}

	// --- cmdb_id ---
	listTesting["CmdbId"] = datamodeltest.TestOptions{
		ValueString: gofakeit.CarModel(),
		SetFunc: func() {
			snapshot.SetAnyCMDBID(listTesting["CmdbId"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetCMDBID(), listTesting["CmdbId"].ValueString)
		},
	}

	// --- os_type ---
	listTesting["OSType"] = datamodeltest.TestOptions{
		ValueString: gofakeit.BankType(),
		SetFunc: func() {
			snapshot.SetAnyOSType(listTesting["OSType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetOSType(), listTesting["OSType"].ValueString)
		},
	}

	// --- hostname ---
	listTesting["Hostname"] = datamodeltest.TestOptions{
		ValueString: gofakeit.DomainName(),
		SetFunc: func() {
			snapshot.SetAnyHostname(listTesting["Hostname"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetHostname(), listTesting["Hostname"].ValueString)
		},
	}

	// --- severity ---
	listTesting["Severity"] = datamodeltest.TestOptions{
		ValueString: gofakeit.BeerHop(),
		SetFunc: func() {
			snapshot.SetAnySeverity(listTesting["Severity"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetSeverity(), listTesting["Severity"].ValueString)
		},
	}

	// --- user_cmdb_name ---
	listTesting["UserCmdbName"] = datamodeltest.TestOptions{
		ValueString: gofakeit.Username(),
		SetFunc: func() {
			snapshot.SetAnyUserCMDBName(listTesting["UserCmdbName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetUserCMDBName(), listTesting["UserCmdbName"].ValueString)
		},
	}

	// --- user_cmdb_id ---
	listTesting["UserCmdbId"] = datamodeltest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			snapshot.SetAnyUserCmdbId(listTesting["UserCmdbId"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, snapshot.GetUserCmdbId(), listTesting["UserCmdbId"].ValueString)
		},
	}

	// --- ip_addresses ---
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

	// --- mac_addresses ---
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
