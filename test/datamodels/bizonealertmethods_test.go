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
		assert.NoError(t, snapshot.SetAnyIPAddresse(ipAddresses))
		assert.NoError(t, snapshot.SetAnyMACAddresse(macAddresses))
		assert.NoError(t, snapshot.SetAnyOS(gofakeit.BankName()))
		assert.NoError(t, snapshot.SetAnyOSType(gofakeit.BankType()))
		assert.NoError(t, snapshot.SetAnyFqdn(gofakeit.DomainName()))
		assert.NoError(t, snapshot.SetAnyDomain(gofakeit.DomainName()))
		assert.NoError(t, snapshot.SetAnyCMDBID(gofakeit.CarModel()))
		assert.NoError(t, snapshot.SetAnyHostname(gofakeit.DomainName()))
		assert.NoError(t, snapshot.SetAnyTitle(gofakeit.AppAuthor()))
		assert.NoError(t, snapshot.SetAnySeverity(gofakeit.Sentence()))
		assert.NoError(t, snapshot.SetAnyUserCmdbId(gofakeit.ID()))
		assert.NoError(t, snapshot.SetAnyUserCMDBName(gofakeit.DomainSuffix()))

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

					if a.GetTitle() != b.GetTitle() {
						return false
					}

					if a.GetSeverity() != b.GetSeverity() {
						return false
					}

					if a.GetUserCmdbId() != b.GetUserCmdbId() {
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
		assert.NoError(t, tag.SetAnyName(gofakeit.EmojiTag()))
		assert.NoError(t, tag.SetAnyColor(gofakeit.Color()))
		assert.NoError(t, tag.SetAnyCreated(gofakeit.Date().String()))
		assert.NoError(t, tag.SetAnyCreatedByUsername(gofakeit.Name()))
		assert.NoError(t, tag.SetAnyCreatedByID(gofakeit.Uint64()))
		assert.NoError(t, tag.SetAnyIsVisibleForCustomer(true))

		tagsExample = append(tagsExample, *tag)
	}

	listTesting["Tags"] = datamodelstest.TestOptions{
		ValueAny: tagsExample,
		SetFunc: func() {
			biZoneIRPAlert.SetTags(tagsExample)
		},
		GetFunc: func() {
			tags, ok := listTesting["Tags"].ValueAny.([]datamodels.BiZoneIRPTag)
			assert.True(t, ok)

			assert.Equal(t, tags, biZoneIRPAlert.GetTags())
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
	assert.NoError(t, dataExample.SetAgent(gofakeit.Uint64()))
	assert.NoError(t, dataExample.SetSeverityID(gofakeit.Uint64()))
	assert.NoError(t, dataExample.SetDesc(gofakeit.Adjective()))
	assert.NoError(t, dataExample.SetEventUid(gofakeit.ID()))
	assert.NoError(t, dataExample.SetJobTitle(gofakeit.Book().Title))
	assert.NoError(t, dataExample.SetFirstSeenTime(gofakeit.Date().String()))
	assert.NoError(t, dataExample.SetLastSeenTime(gofakeit.Date().String()))
	assert.NoError(t, dataExample.SetMetadataProductName(gofakeit.PetName()))
	assert.NoError(t, dataExample.SetSourceIP(gofakeit.IPv4Address()))
	assert.NoError(t, dataExample.SetTargetIP(gofakeit.IPv4Address()))
	assert.NoError(t, dataExample.SetUnmappedHiveAlertID(gofakeit.ID()))
	assert.NoError(t, dataExample.SetUnmappedSensorIP(gofakeit.IPv4Address()))
	assert.NoError(t, dataExample.SetUnmappedSensorName(gofakeit.PetName()))

	// --- tags ---
	dataTagsSize := 11
	dataTags := make([]string, 0, dataTagsSize)
	for range dataTagsSize {
		dataTags = append(dataTags, gofakeit.Adjective())
	}
	assert.NoError(t, dataExample.SetTags(dataTags))

	// --- unmapped_dst_endpoint_array ---
	dataUnmappedDstEndpointArraySize := 12
	dataUnmappedDstEndpointArray := make([]string, 0, dataUnmappedDstEndpointArraySize)
	for range dataUnmappedDstEndpointArraySize {
		dataUnmappedDstEndpointArray = append(dataUnmappedDstEndpointArray, gofakeit.FarmAnimal())
	}
	assert.NoError(t, dataExample.SetUnmappedDstEndpointArray(dataUnmappedDstEndpointArray))

	// --- unmapped_home_endpoint_array ---
	dataUnmappedHomeEndpointArraySize := 13
	dataUnmappedHomeEndpointArray := make([]string, 0, dataUnmappedHomeEndpointArraySize)
	for range dataUnmappedHomeEndpointArraySize {
		dataUnmappedHomeEndpointArray = append(dataUnmappedHomeEndpointArray, gofakeit.FarmAnimal())
	}
	assert.NoError(t, dataExample.SetUnmappedDstEndpointArray(dataUnmappedHomeEndpointArray))

	// --- detection_pattern ---
	dataDetectionPatternSize := 13
	dataDetectionPattern := make([]uint64, 0, dataDetectionPatternSize)
	for range dataDetectionPatternSize {
		dataDetectionPattern = append(dataDetectionPattern, gofakeit.Uint64())
	}
	assert.NoError(t, dataExample.SetUnmappedDstEndpointArray(dataUnmappedHomeEndpointArray))

	// --- unmapped_agent_array ---
	dataUnmappedAgentArraySize := 13
	dataUnmappedAgentArray := make([]uint64, 0, dataUnmappedAgentArraySize)
	for range dataUnmappedAgentArraySize {
		dataUnmappedAgentArray = append(dataUnmappedAgentArray, gofakeit.Uint64())
	}
	assert.NoError(t, dataExample.SetUnmappedAgentArrayn(dataUnmappedAgentArray))

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
			biZoneIRPAlert.SetAnyID(listTesting["IDNum"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPAlert.GetID(), listTesting["IDNum"].ValueAny)
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
