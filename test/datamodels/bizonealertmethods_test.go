package datamodels_test

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneIRPAlertMethods(t *testing.T) {
	var (
		size int
	)

	biZoneIRPAlert := &datamodels.VerifiedBiZoneIRPAlert{}
	listTesting := map[string]datamodelstest.TestOptions{}

	// ---- Snapshots ----
	size = 13
	ipAddresses := make([]string, 0, size)
	for range size {
		ipAddresses = append(ipAddresses, gofakeit.IPv4Address())
	}
	macAddresses := make([]string, 0, size)
	for range size {
		macAddresses = append(macAddresses, gofakeit.MacAddress())
	}

	snapshotsExample := make([]datamodels.BiZoneIRPSnapshot, 0, size)
	for range size {
		snapshot := datamodels.NewBiZoneIRPSnapshot()
		snapshot.SetAnyIPAddresse(ipAddresses)
		snapshot.SetAnyMACAddresse(macAddresses)
		snapshot.SetAnyOS(gofakeit.BankName())
		snapshot.SetAnyOSType(gofakeit.BankType())
		snapshot.SetAnyFqdn(gofakeit.DomainName())
		snapshot.SetAnyDomain(gofakeit.DomainName())
		snapshot.SetAnyCMDBID(gofakeit.CarModel())
		snapshot.SetAnyHostname(gofakeit.DomainName())
		snapshot.SetAnyUserCMDBName(gofakeit.DomainSuffix())

		snapshotsExample = append(snapshotsExample, *snapshot)
	}

	listTesting["Snapshots"] = datamodelstest.TestOptions{
		ValueAny: snapshotsExample,
		SetFunc: func() {
			biZoneIRPAlert.SetSnapshots(snapshotsExample)
		},
		GetFunc: func() {
			snapshots, ok := listTesting["Snapshots"].ValueAny.([]datamodels.BiZoneIRPSnapshot)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				snapshots,
				biZoneIRPAlert.GetSnapshots(),
				func(a, b datamodels.BiZoneIRPSnapshot) bool {
					if !slices.Equal(a.IPAddresses, b.IPAddresses) {
						return false
					}

					if !slices.Equal(a.MACAddresses, b.MACAddresses) {
						return false
					}

					if a.GetOS() != b.GetOS() {
						return false
					}

					if a.GetOSType() != b.GetOSType() {
						return false
					}

					if a.GetFqdn() != b.GetFqdn() {
						return false
					}

					if a.GetDomain() != b.GetDomain() {
						return false
					}

					if a.GetCMDBID() != b.GetCMDBID() {
						return false
					}
					if a.GetHostname() != b.GetHostname() {
						return false
					}
					return a.GetUserCMDBName() == b.GetUserCMDBName()
				}))
		},
	}

	// ---- Tags ----
	size = 14
	tagsExample := make([]datamodels.BiZoneIRPTag, 0, size)
	for range size {
		tag := datamodels.NewBiZoneIRPTag()
		tag.SetAnyName(gofakeit.EmojiTag())
		tag.SetAnyColor(gofakeit.Color())

		tagsExample = append(tagsExample, *tag)
	}

	listTesting["Tags"] = datamodelstest.TestOptions{
		ValueAny: snapshotsExample,
		SetFunc: func() {
			biZoneIRPAlert.SetSnapshots(snapshotsExample)
		},
		GetFunc: func() {
			snapshots, ok := listTesting["Tags"].ValueAny.([]datamodels.BiZoneIRPSnapshot)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				snapshots,
				biZoneIRPAlert.GetSnapshots(),
				func(a, b datamodels.BiZoneIRPSnapshot) bool {
					if !slices.Equal(a.IPAddresses, b.IPAddresses) {
						return false
					}

					if !slices.Equal(a.MACAddresses, b.MACAddresses) {
						return false
					}

					if a.GetOS() != b.GetOS() {
						return false
					}

					if a.GetOSType() != b.GetOSType() {
						return false
					}

					if a.GetFqdn() != b.GetFqdn() {
						return false
					}

					if a.GetDomain() != b.GetDomain() {
						return false
					}

					if a.GetCMDBID() != b.GetCMDBID() {
						return false
					}
					if a.GetHostname() != b.GetHostname() {
						return false
					}
					return a.GetUserCMDBName() == b.GetUserCMDBName()
				}))
		},
	}

	// ---- Data ----
	size = 11
	allIPHomes := make([]string, 0, size)
	for range size {
		allIPHomes = append(allIPHomes, gofakeit.IPv4Address())
	}

	snortIds := make([]uint64, 0, size+2)
	for range size {
		snortIds = append(snortIds, gofakeit.Uint64())
	}

	allSensors := make([]uint64, 0, size+1)
	for range size {
		allSensors = append(allSensors, gofakeit.Uint64())
	}

	dataExample := datamodels.NewBiZoneIRPData()
	dataExample.SetAnyIPHome(allIPHomes)
	dataExample.SetAnySnortSid(snortIds)
	dataExample.SetAnyAllSensor(allSensors)
	dataExample.SetAnyID(gofakeit.ID())
	dataExample.SetAnyTitle(gofakeit.BookTitle())
	dataExample.SetAnyIPHome(gofakeit.IPv4Address())
	dataExample.SetAnyURLFTP(gofakeit.URL())
	dataExample.SetAnyIPExter(gofakeit.IPv4Address())
	dataExample.SetAnyURLHTTP(gofakeit.URL())
	dataExample.SetAnyEventType(gofakeit.BankType())
	dataExample.SetAnyURLArkime(gofakeit.URL())
	dataExample.SetAnySensor(gofakeit.Uint64())
	dataExample.SetAnyAllIPExt(gofakeit.Uint64())
	dataExample.SetAnyResponseTeam(gofakeit.Uint64())

	listTesting["Data"] = datamodelstest.TestOptions{
		ValueAny: *dataExample,
		SetFunc: func() {
			biZoneIRPAlert.SetData(*dataExample)
		},
		GetFunc: func() {
			data, ok := listTesting["Data"].ValueAny.(datamodels.BiZoneIRPData)
			assert.True(t, ok)
			assert.Equal(t, data, *biZoneIRPAlert.GetData())
		},
	}

	// ---- SpecialUUID ----
	listTesting["SpecialUUID"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.UUID(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyUUID(listTesting["SpecialUUID"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetUUID(), listTesting["SpecialUUID"].ValueAny)
		},
	}

	// ---- UUID ----
	listTesting["UUID"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.UUID(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyUUID(listTesting["UUID"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetUUID(), listTesting["UUID"].ValueAny)
		},
	}

	// ---- Title ----
	listTesting["Title"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.BookTitle(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyTitle(listTesting["Title"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetTitle(), listTesting["Title"].ValueAny)
		},
	}

	// ---- Severity ----
	listTesting["Severity"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.BookGenre(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnySeverity(listTesting["Severity"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetSeverity(), listTesting["Severity"].ValueAny)
		},
	}

	// ---- ExternalID ----
	listTesting["ExternalID"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.ID(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyExternalID(listTesting["ExternalID"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetExternalID(), listTesting["ExternalID"].ValueAny)
		},
	}

	// ---- Confidence ----
	listTesting["Confidence"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.BookAuthor(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyConfidence(listTesting["Confidence"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetConfidence(), listTesting["Confidence"].ValueAny)
		},
	}

	// ---- Description ----
	listTesting["Description"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Dessert(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyDescription(listTesting["Description"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetDescription(), listTesting["Description"].ValueAny)
		},
	}

	// ---- CreatedTime ----
	listTesting["CreatedTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyCreatedTime(listTesting["CreatedTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetCreatedTime(), listTesting["CreatedTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- UpdatedTime ----
	listTesting["UpdatedTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyUpdatedTime(listTesting["UpdatedTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetUpdatedTime(), listTesting["UpdatedTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- PlatformType ----
	listTesting["PlatformType"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.AnimalType(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyPlatformType(listTesting["PlatformType"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetPlatformType(), listTesting["PlatformType"].ValueAny)
		},
	}

	// ---- EventStartTime ----
	listTesting["EventStartTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyEventStartTime(listTesting["EventStartTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetEventStartTime(), listTesting["EventStartTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- EventEndTime ----
	listTesting["EventEndTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyEventEndTime(listTesting["EventEndTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetEventEndTime(), listTesting["EventEndTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- DetectionRule ----
	listTesting["DetectionRule"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.NounDeterminer(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyDetectionRule(listTesting["DetectionRule"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetDetectionRule(), listTesting["DetectionRule"].ValueAny)
		},
	}

	// ---- CustomerSystem ----
	listTesting["CustomerSystem"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.AchAccount(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyCustomerSystem(listTesting["CustomerSystem"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetCustomerSystem(), listTesting["CustomerSystem"].ValueAny)
		},
	}

	// ---- Recommendations ----
	listTesting["Recommendations"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Animal(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyRecommendations(listTesting["Recommendations"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetRecommendations(), listTesting["Recommendations"].ValueAny)
		},
	}

	// ---- PlatformHostname ----
	listTesting["PlatformHostname"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.MinecraftMobHostile(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyPlatformHostname(listTesting["PlatformHostname"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetPlatformHostname(), listTesting["PlatformHostname"].ValueAny)
		},
	}

	// ---- FirstDetectionTime ----
	listTesting["FirstDetectionTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyFirstDetectionTime(listTesting["FirstDetectionTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetFirstDetectionTime(), listTesting["FirstDetectionTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- LastDetectionTime ----
	listTesting["LastDetectionTime"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPAlert.SetAnyLastDetectionTime(listTesting["LastDetectionTime"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetLastDetectionTime(), listTesting["LastDetectionTime"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- IDNum ----
	listTesting["IDNum"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Uint64(),
		SetFunc: func() {
			biZoneIRPAlert.SetAnyIDNum(listTesting["IDNum"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetIDNum(), listTesting["IDNum"].ValueAny)
		},
	}

	/*
		// ----  ----
		listTesting[""] = datamodelstest.TestOptions{
			ValueAny: gofakeit.TimeZoneRegion(),
			SetFunc: func() {
				biZoneIRPAlert.SetAny(listTesting[""].ValueAny)
			},
			GetFunc: func() {
				assert.Equal(t, biZoneIRPAlert.Get(), listTesting[""].ValueAny)
			},
		}


		listTesting[""] = datamodelstest.TestOptions{
			ValueString: ,
			SetFunc: func() {
				biZoneIRPCase.(listTesting[""].ValueString)
			},
			GetFunc: func() {
				assert.Equal(t, biZoneIRPCase.(), listTesting[""].ValueString)
			},
		}
	*/

	var num int
	for k, v := range listTesting {
		num++
		t.Run(fmt.Sprintf("Test %d. Field %s", num, k), func(t *testing.T) {
			v.SetFunc()
			v.GetFunc()
		})
	}
}
