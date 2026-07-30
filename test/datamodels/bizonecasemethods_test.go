package datamodels_test

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneIRPCaseMethods(t *testing.T) {
	var (
		size      int
		keyString string
		valueAny  any
	)

	biZoneIRPCase := &datamodels.VerifiedIRPBiZoneCase{}
	listTesting := map[string]datamodelstest.TestOptions{}

	// ---- GossopkaErrors ----
	size = 10
	gossopkaErrorsExample := make(map[string]any, size)
	for range size {
		gossopkaErrorsExample[gofakeit.Error().Error()] = gofakeit.ErrorDatabase()
	}

	listTesting["GossopkaErrors"] = datamodelstest.TestOptions{
		ValueMapAny: gossopkaErrorsExample,
		SetFunc: func() {
			biZoneIRPCase.SetGossopkaErrors(listTesting["GossopkaErrors"].ValueMapAny)

			keyString = gofakeit.Error().Error()
			valueAny = gofakeit.ErrorDatabase()
			assert.True(t, biZoneIRPCase.SetAnyGossopkaErrors(keyString, valueAny))

			listTesting["GossopkaErrors"].ValueMapAny[keyString] = valueAny
		},
		GetFunc: func() {
			assert.True(t, maps.EqualFunc(
				listTesting["GossopkaErrors"].ValueMapAny,
				biZoneIRPCase.GetGossopkaErrors(),
				func(a1 any, a2 any) bool {
					v1, ok := a1.(error)
					if !ok {
						return false
					}

					v2, ok := a2.(error)
					if !ok {
						return false
					}

					return strings.EqualFold(v1.Error(), v2.Error())
				}))
		},
	}

	// ---- ObservedIocs ----
	size = 10
	observedIocsExample := make([]datamodels.BiZoneIRPObservedIocs, 0, size)

	for range size {
		observedIoc := datamodels.NewBiZoneIRPObservedIocs()
		observedIoc.SetCategories(gofakeit.NiceColors())
		observedIoc.SetAnyIoc(gofakeit.Color())
		observedIoc.SetAnyIocType(gofakeit.CarType())

		observedIocsExample = append(observedIocsExample, *observedIoc)
	}

	listTesting["ObservedIocs"] = datamodelstest.TestOptions{
		ValueAny: observedIocsExample,
		SetFunc: func() {
			biZoneIRPCase.SetObservedIocs(observedIocsExample)
		},
		GetFunc: func() {
			observedIocs, ok := listTesting["ObservedIocs"].ValueAny.([]datamodels.BiZoneIRPObservedIocs)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				observedIocs,
				biZoneIRPCase.GetObservedIocs(),
				func(a, b datamodels.BiZoneIRPObservedIocs) bool {
					if a.GetIoc() != b.GetIoc() {
						return false
					}

					if a.GetIocType() != b.GetIocType() {
						return false
					}

					return slices.Equal(a.GetCategories(), b.GetCategories())
				}))
		},
	}

	// ---- SecondaryCategoryRef ----
	size = 13
	secondaryCategoryRefExample := make([]datamodels.BiZoneIRPTypeRef, 0, size)
	for range size {
		secondaryCategoryRef := datamodels.NewBiZoneIRPTypeRef()
		secondaryCategoryRef.SetAnyId(gofakeit.ID())
		secondaryCategoryRef.SetTitle(gofakeit.JobTitle())
		secondaryCategoryRef.SetDescription(gofakeit.JobDescriptor())

		secondaryCategoryRefExample = append(secondaryCategoryRefExample, *secondaryCategoryRef)
	}

	listTesting["SecondaryCategoryRef"] = datamodelstest.TestOptions{
		ValueAny: secondaryCategoryRefExample,
		SetFunc: func() {
			biZoneIRPCase.SetSecondaryCategoryRef(secondaryCategoryRefExample)
		},
		GetFunc: func() {
			secondaryCategoryRefs, ok := listTesting["SecondaryCategoryRef"].ValueAny.([]datamodels.BiZoneIRPTypeRef)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				secondaryCategoryRefs,
				biZoneIRPCase.GetSecondaryCategoryRef(),
				func(a, b datamodels.BiZoneIRPTypeRef) bool {
					if a.GetId() != b.GetId() {
						return false
					}

					if a.GetTitle() != b.GetTitle() {
						return false
					}

					return a.GetDescription() == b.GetDescription()
				}))
		},
	}

	// ---- Watcher ----
	size = 13
	watchersExample := make([]datamodels.BiZoneIRPWatcher, 0, size)
	for range size {
		watcherExample := datamodels.NewBiZoneIRPWatcher()
		watcherExample.SetAnyId(gofakeit.ID())
		watcherExample.SetUsername(gofakeit.Name())
		watcherExample.SetFirstName(gofakeit.FirstName())
		watcherExample.SetLastName(gofakeit.LastName())
		watcherExample.SetEmail(gofakeit.Email())
		watcherExample.SetPatronimic(gofakeit.HipsterParagraph())
		watcherExample.SetIsActive(gofakeit.Bool())

		watchersExample = append(watchersExample, *watcherExample)
	}

	listTesting["Watchers"] = datamodelstest.TestOptions{
		ValueAny: watchersExample,
		SetFunc: func() {
			biZoneIRPCase.SetWatchers(watchersExample)
		},
		GetFunc: func() {
			watchers, ok := listTesting["Watchers"].ValueAny.([]datamodels.BiZoneIRPWatcher)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				watchers,
				biZoneIRPCase.GetWatchers(),
				func(a, b datamodels.BiZoneIRPWatcher) bool {
					if a.GetId() != b.GetId() {
						return false
					}

					if a.GetUsername() != b.GetUsername() {
						return false
					}

					if a.GetFirstName() != b.GetFirstName() {
						return false
					}

					if a.GetLastName() != b.GetLastName() {
						return false
					}

					if a.GetEmail() != b.GetEmail() {
						return false
					}

					if a.GetPatronimic() != b.GetPatronimic() {
						return false
					}

					return a.GetIsActive() == b.GetIsActive()
				}))
		},
	}

	// ---- ActivityDetected ----
	size = 12
	activityDetectedExp := make([]any, 0, size)
	for range size {
		activityDetectedExp = append(activityDetectedExp, DataStorage[string]{Data: gofakeit.APIUserAgent()})
	}

	listTesting["ActivityDetected"] = datamodelstest.TestOptions{
		ValueAny: activityDetectedExp,
		SetFunc: func() {
			biZoneIRPCase.SetActivityDetected(activityDetectedExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetActivityDetected(), listTesting["ActivityDetected"].ValueAny)
		},
	}

	// ---- Attachments ----
	size = 15
	attachmentsExp := make([]any, 0, size)
	for range size {
		attachmentsExp = append(attachmentsExp, DataStorage[string]{Data: gofakeit.CarModel()})
	}

	listTesting["Attachments"] = datamodelstest.TestOptions{
		ValueAny: attachmentsExp,
		SetFunc: func() {
			biZoneIRPCase.SetAttachments(attachmentsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAttachments(), listTesting["Attachments"].ValueAny)
		},
	}

	// ---- Badges ----
	size = 11
	badgesExp := make([]any, 0, size)
	for range size {
		badgesExp = append(badgesExp, DataStorage[string]{Data: gofakeit.City()})
	}

	listTesting["Badges"] = datamodelstest.TestOptions{
		ValueAny: badgesExp,
		SetFunc: func() {
			biZoneIRPCase.SetBadges(badgesExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetBadges(), listTesting["Badges"].ValueAny)
		},
	}

	// ---- Emls ----
	size = 12
	emlsExp := make([]any, 0, size)
	for range size {
		emlsExp = append(emlsExp, DataStorage[string]{Data: gofakeit.TimeZone()})
	}

	listTesting["Emls"] = datamodelstest.TestOptions{
		ValueAny: emlsExp,
		SetFunc: func() {
			biZoneIRPCase.SetEmls(emlsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetEmls(), listTesting["Emls"].ValueAny)
		},
	}

	// ---- Slas ----
	size = 10
	slasExp := make([]any, 0, size)
	for range size {
		slasExp = append(slasExp, DataStorage[string]{Data: gofakeit.NounCommon()})
	}

	listTesting["Slas"] = datamodelstest.TestOptions{
		ValueAny: slasExp,
		SetFunc: func() {
			biZoneIRPCase.SetSlas(slasExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSlas(), listTesting["Slas"].ValueAny)
		},
	}

	// ---- Tags ----
	size = 13
	tagsExp := make([]any, 0, size)
	for range size {
		tagsExp = append(tagsExp, DataStorage[string]{Data: gofakeit.EmojiTag()})
	}

	listTesting["Tags"] = datamodelstest.TestOptions{
		ValueAny: tagsExp,
		SetFunc: func() {
			biZoneIRPCase.SetTags(tagsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetTags(), listTesting["Tags"].ValueAny)
		},
	}

	// ---- KeyField ----
	size = 12
	keyFieldExp := make([]any, 0, size)
	for range size {
		keyFieldExp = append(keyFieldExp, DataStorage[string]{Data: gofakeit.AnimalType()})
	}

	listTesting["KeyField"] = datamodelstest.TestOptions{
		ValueAny: keyFieldExp,
		SetFunc: func() {
			biZoneIRPCase.SetKeyField(keyFieldExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetKeyField(), listTesting["KeyField"].ValueAny)
		},
	}

	// ---- MITRECOV ----
	size = 11
	mitreCovExp := datamodels.NewBiZoneMITRECOV()
	for range size {
		mitreCovExp.AddAnyPersistence(gofakeit.AppAuthor())
	}

	listTesting["MITRECOV"] = datamodelstest.TestOptions{
		ValueAny: *mitreCovExp,
		SetFunc: func() {
			biZoneIRPCase.SetMITRECOV(*mitreCovExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetMITRECOV(), listTesting["MITRECOV"].ValueAny)
		},
	}

	// ---- DetectionRules ----
	size = 10
	detectionRulesExp := make([]string, 0, size)
	for range size {
		detectionRulesExp = append(detectionRulesExp, gofakeit.Dessert())
	}

	listTesting["DetectionRules"] = datamodelstest.TestOptions{
		ValueSliceString: detectionRulesExp,
		SetFunc: func() {
			biZoneIRPCase.SetDetectionRules(detectionRulesExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetDetectionRules(), listTesting["DetectionRules"].ValueSliceString)
		},
	}

	// ---- SecondaryCategory ----
	size = 15
	secondaryCategoryExp := make([]string, 0, size)
	for range size {
		secondaryCategoryExp = append(secondaryCategoryExp, gofakeit.AchRouting())
	}

	listTesting["SecondaryCategory"] = datamodelstest.TestOptions{
		ValueSliceString: secondaryCategoryExp,
		SetFunc: func() {
			biZoneIRPCase.SetSecondaryCategories(secondaryCategoryExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSecondaryCategory(), listTesting["SecondaryCategory"].ValueSliceString)
		},
	}

	// ---- PlatformHostname ----
	size = 15
	platformHostnameExp := make([]string, 0, size)
	for range size {
		platformHostnameExp = append(platformHostnameExp, gofakeit.AppName())
	}

	listTesting["PlatformHostname"] = datamodelstest.TestOptions{
		ValueSliceString: platformHostnameExp,
		SetFunc: func() {
			biZoneIRPCase.SetPlatformHostname(platformHostnameExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetPlatformHostname(), listTesting["PlatformHostname"].ValueSliceString)
		},
	}

	// ---- WatchersId ----
	size = 15
	watchersIdExp := make([]uint64, 0, size)
	for range size {
		watchersIdExp = append(watchersIdExp, gofakeit.Uint64())
	}

	listTesting["WatchersId"] = datamodelstest.TestOptions{
		ValueSliceUint64: watchersIdExp,
		SetFunc: func() {
			biZoneIRPCase.SetWatchersId(watchersIdExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetWatchersId(), listTesting["WatchersId"].ValueSliceUint64)
		},
	}

	// ---- Assignee ----
	listTesting["Assignee"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.FarmAnimal(),
		SetFunc: func() {
			biZoneIRPCase.SetAssignee(listTesting["Assignee"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAssignee(), listTesting["Assignee"].ValueAny)
		},
	}

	// ---- AssigneeDisplayName ----
	listTesting["AssigneeDisplayName"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.School(),
		SetFunc: func() {
			biZoneIRPCase.SetAssigneeDisplayName(listTesting["AssigneeDisplayName"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAssigneeDisplayName(), listTesting["AssigneeDisplayName"].ValueAny)
		},
	}

	// ---- Service ----
	listTesting["Service"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.AirlineSeat(),
		SetFunc: func() {
			biZoneIRPCase.SetService(listTesting["Service"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetService(), listTesting["Service"].ValueAny)
		},
	}

	// ---- ResolutionDate ----
	listTesting["ResolutionDate"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Dinner(),
		SetFunc: func() {
			biZoneIRPCase.SetResolutionDate(listTesting["ResolutionDate"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetResolutionDate(), listTesting["ResolutionDate"].ValueAny)
		},
	}

	// ---- ResolutionName ----
	listTesting["ResolutionName"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Question(),
		SetFunc: func() {
			biZoneIRPCase.SetResolutionName(listTesting["ResolutionName"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetResolutionName(), listTesting["ResolutionName"].ValueAny)
		},
	}

	// ---- ResolutionNameRef ----
	listTesting["ResolutionNameRef"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.PronounReflective(),
		SetFunc: func() {
			biZoneIRPCase.SetResolutionNameRef(listTesting["ResolutionNameRef"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetResolutionNameRef(), listTesting["ResolutionNameRef"].ValueAny)
		},
	}

	// ---- GossopkaId ----
	listTesting["GossopkaId"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.GlobalFaker.ID(),
		SetFunc: func() {
			biZoneIRPCase.SetGossopkaId(listTesting["GossopkaId"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGossopkaId(), listTesting["GossopkaId"].ValueAny)
		},
	}

	// ---- GossopkaSendTime ----
	listTesting["GossopkaSendTime"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.TimeZoneFull(),
		SetFunc: func() {
			biZoneIRPCase.SetGossopkaSendTime(listTesting["GossopkaSendTime"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGossopkaSendTime(), listTesting["GossopkaSendTime"].ValueAny)
		},
	}

	// ---- GtiId ----
	listTesting["GtiId"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.ID(),
		SetFunc: func() {
			biZoneIRPCase.SetGtiId(listTesting["GtiId"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGtiId(), listTesting["GtiId"].ValueAny)
		},
	}

	// ---- GtiSendTime ----
	listTesting["GtiSendTime"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.TimeZoneRegion(),
		SetFunc: func() {
			biZoneIRPCase.SetGtiSendTime(listTesting["GtiSendTime"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGtiSendTime(), listTesting["GtiSendTime"].ValueAny)
		},
	}

	// ---- CustomerStarRating ----
	listTesting["CustomerStarRating"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.AdjectiveDemonstrative(),
		SetFunc: func() {
			biZoneIRPCase.SetCustomerStarRating(listTesting["CustomerStarRating"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetCustomerStarRating(), listTesting["CustomerStarRating"].ValueAny)
		},
	}

	// ---- FirstSentNotificationTime ----
	listTesting["FirstSentNotificationTime"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.AnimalType(),
		SetFunc: func() {
			biZoneIRPCase.SetFirstSentNotificationTime(listTesting["FirstSentNotificationTime"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetFirstSentNotificationTime(), listTesting["FirstSentNotificationTime"].ValueAny)
		},
	}

	// ---- ResponseTeam ----
	listTesting["ResponseTeam"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.Zip(),
		SetFunc: func() {
			biZoneIRPCase.SetResponseTeam(listTesting["ResponseTeam"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetResponseTeam(), listTesting["ResponseTeam"].ValueAny)
		},
	}

	// ---- GossopkaKeyLink ----
	listTesting["GossopkaKeyLink"] = datamodelstest.TestOptions{
		ValueAny: gofakeit.FarmAnimal(),
		SetFunc: func() {
			biZoneIRPCase.SetGossopkaKeyLink(listTesting["GossopkaKeyLink"].ValueAny)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGossopkaKeyLink(), listTesting["GossopkaKeyLink"].ValueAny)
		},
	}

	// ---- IssueTypeRef ----
	IssueTypeRef := datamodels.NewBiZoneIRPTypeRef()
	IssueTypeRef.SetAnyId(gofakeit.ID())
	IssueTypeRef.SetAnyTitle(gofakeit.BookTitle())
	IssueTypeRef.SetAnyDescription(gofakeit.BookGenre())

	listTesting["IssueTypeRef"] = datamodelstest.TestOptions{
		ValueAny: *IssueTypeRef,
		SetFunc: func() {
			biZoneIRPCase.SetIssueTypeRef(*IssueTypeRef)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetIssueTypeRef(), listTesting["IssueTypeRef"].ValueAny)
		},
	}

	// ---- PrimaryCategoryRef ----
	primaryCategoryRef := datamodels.NewBiZoneIRPTypeRef()
	primaryCategoryRef.SetAnyId(gofakeit.ID())
	primaryCategoryRef.SetAnyTitle(gofakeit.BookTitle())
	primaryCategoryRef.SetAnyDescription(gofakeit.BookGenre())

	listTesting["PrimaryCategoryRef"] = datamodelstest.TestOptions{
		ValueAny: *primaryCategoryRef,
		SetFunc: func() {
			biZoneIRPCase.SetPrimaryCategoryRef(*primaryCategoryRef)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetPrimaryCategoryRef(), listTesting["PrimaryCategoryRef"].ValueAny)
		},
	}

	// ---- AttackType ----
	listTesting["AttackType"] = datamodelstest.TestOptions{
		ValueString: gofakeit.CarType(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyAttackType(listTesting["AttackType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAttackType(), listTesting["AttackType"].ValueString)
		},
	}

	// ---- Created ----
	listTesting["Created"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPCase.SetAnyCreated(listTesting["Created"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetCreated(), listTesting["Created"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- CreatorDisplayName ----
	listTesting["CreatorDisplayName"] = datamodelstest.TestOptions{
		ValueString: gofakeit.BeerName(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyCreatorDisplayName(listTesting["CreatorDisplayName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetCreatorDisplayName(), listTesting["CreatorDisplayName"].ValueString)
		},
	}

	// ---- DetectionDate ----
	listTesting["DetectionDate"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPCase.SetAnyDetectionDate(listTesting["DetectionDate"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetDetectionDate(), listTesting["DetectionDate"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- Key ----
	listTesting["Key"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyKey(listTesting["Key"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetKey(), listTesting["Key"].ValueString)
		},
	}

	// ---- IssueType ----
	listTesting["IssueType"] = datamodelstest.TestOptions{
		ValueString: gofakeit.BankType(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyIssueType(listTesting["IssueType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetIssueType(), listTesting["IssueType"].ValueString)
		},
	}

	// ---- Priority ----
	listTesting["Priority"] = datamodelstest.TestOptions{
		ValueString: gofakeit.TimeZoneRegion(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyPriority(listTesting["Priority"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetPriority(), listTesting["Priority"].ValueString)
		},
	}

	// ---- Summary ----
	listTesting["Summary"] = datamodelstest.TestOptions{
		ValueString: gofakeit.StreetSuffix(),
		SetFunc: func() {
			biZoneIRPCase.SetAnySummary(listTesting["Summary"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSummary(), listTesting["Summary"].ValueString)
		},
	}

	// ---- Description ----
	listTesting["Description"] = datamodelstest.TestOptions{
		ValueString: gofakeit.BookTitle(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyDescription(listTesting["Description"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetDescription(), listTesting["Description"].ValueString)
		},
	}

	// ---- RenderedDescription ----
	listTesting["RenderedDescription"] = datamodelstest.TestOptions{
		ValueString: gofakeit.AnimalType(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyRenderedDescription(listTesting["RenderedDescription"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetRenderedDescription(), listTesting["RenderedDescription"].ValueString)
		},
	}

	// ---- Status ----
	listTesting["Status"] = datamodelstest.TestOptions{
		ValueString: gofakeit.SafeColor(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyStatus(listTesting["Status"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetStatus(), listTesting["Status"].ValueString)
		},
	}

	// ---- StatusDescription ----
	listTesting["StatusDescription"] = datamodelstest.TestOptions{
		ValueString: gofakeit.SafeColor(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyStatusDescription(listTesting["StatusDescription"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetStatusDescription(), listTesting["StatusDescription"].ValueString)
		},
	}

	// ---- ReporterDisplayName ----
	listTesting["ReporterDisplayName"] = datamodelstest.TestOptions{
		ValueString: gofakeit.NamePrefix(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyReporterDisplayName(listTesting["ReporterDisplayName"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetReporterDisplayName(), listTesting["ReporterDisplayName"].ValueString)
		},
	}

	// ---- ReporterEmailAddress ----
	listTesting["ReporterEmailAddress"] = datamodelstest.TestOptions{
		ValueString: gofakeit.Email(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyReporterEmailAddress(listTesting["ReporterEmailAddress"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetReporterEmailAddress(), listTesting["ReporterEmailAddress"].ValueString)
		},
	}

	// ---- PrimaryCategory ----
	listTesting["PrimaryCategory"] = datamodelstest.TestOptions{
		ValueString: gofakeit.CarTransmissionType(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyPrimaryCategory(listTesting["PrimaryCategory"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetPrimaryCategory(), listTesting["PrimaryCategory"].ValueString)
		},
	}

	// ---- Updated ----
	listTesting["Updated"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPCase.SetAnyUpdated(listTesting["Updated"].ValueTime.Format(time.UnixDate)))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetUpdated(), listTesting["Updated"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- Timestamp ----
	listTesting["Timestamp"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPCase.SetAnyTimestamp(listTesting["Timestamp"].ValueTime.Format(time.RFC1123Z)))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetTimestamp(), listTesting["Timestamp"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- ResolutionDetailed ----
	listTesting["ResolutionDetailed"] = datamodelstest.TestOptions{
		ValueString: gofakeit.Gamertag(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyResolutionDetailed(listTesting["ResolutionDetailed"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetResolutionDetailed(), listTesting["ResolutionDetailed"].ValueString)
		},
	}

	// ---- PlatformURL ----
	listTesting["PlatformURL"] = datamodelstest.TestOptions{
		ValueString: gofakeit.URL(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyPlatformURL(listTesting["PlatformURL"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetPlatformURL(), listTesting["PlatformURL"].ValueString)
		},
	}

	// ---- CustomerStarRatingComment ----
	listTesting["CustomerStarRatingComment"] = datamodelstest.TestOptions{
		ValueString: gofakeit.AirlineRecordLocator(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyCustomerStarRatingComment(listTesting["CustomerStarRatingComment"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetCustomerStarRatingComment(), listTesting["CustomerStarRatingComment"].ValueString)
		},
	}

	// ---- Recommendations ----
	listTesting["Recommendations"] = datamodelstest.TestOptions{
		ValueString: gofakeit.TimeZoneRegion(),
		SetFunc: func() {
			biZoneIRPCase.SetRecommendations(listTesting["Recommendations"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetRecommendations(), listTesting["Recommendations"].ValueString)
		},
	}

	// ---- ParsedActivityDetected ----
	listTesting["ParsedActivityDetected"] = datamodelstest.TestOptions{
		ValueString: gofakeit.VerbAction(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyParsedActivityDetected(listTesting["ParsedActivityDetected"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetParsedActivityDetected(), listTesting["ParsedActivityDetected"].ValueString)
		},
	}

	// ---- AffectedService ----
	listTesting["AffectedService"] = datamodelstest.TestOptions{
		ValueString: gofakeit.StreetNumber(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyAffectedService(listTesting["AffectedService"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAffectedService(), listTesting["AffectedService"].ValueString)
		},
	}

	// ---- FakeAsPath ----
	listTesting["FakeAsPath"] = datamodelstest.TestOptions{
		ValueString: gofakeit.Paragraph(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyFakeAsPath(listTesting["FakeAsPath"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetFakeAsPath(), listTesting["FakeAsPath"].ValueString)
		},
	}

	// ---- FakePrefix ----
	listTesting["FakePrefix"] = datamodelstest.TestOptions{
		ValueString: gofakeit.NamePrefix(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyFakePrefix(listTesting["FakePrefix"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetFakePrefix(), listTesting["FakePrefix"].ValueString)
		},
	}

	// ---- FpType ----
	listTesting["FpType"] = datamodelstest.TestOptions{
		ValueString: gofakeit.TimeZoneRegion(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyFpType(listTesting["FpType"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetFpType(), listTesting["FpType"].ValueString)
		},
	}

	// ---- LookingGlass ----
	listTesting["LookingGlass"] = datamodelstest.TestOptions{
		ValueString: gofakeit.TimeZoneRegion(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyLookingGlass(listTesting["LookingGlass"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetLookingGlass(), listTesting["LookingGlass"].ValueString)
		},
	}

	// ---- SpamRecipients ----
	listTesting["SpamRecipients"] = datamodelstest.TestOptions{
		ValueString: gofakeit.BitcoinPrivateKey(),
		SetFunc: func() {
			biZoneIRPCase.SetAnySpamRecipients(listTesting["SpamRecipients"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSpamRecipients(), listTesting["SpamRecipients"].ValueString)
		},
	}

	// ---- TLP ----
	listTesting["TLP"] = datamodelstest.TestOptions{
		ValueString: gofakeit.AdverbFrequencyDefinite(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyTLP(listTesting["TLP"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetTLP(), listTesting["TLP"].ValueString)
		},
	}

	// ---- UsualPrefix ----
	listTesting["UsualPrefix"] = datamodelstest.TestOptions{
		ValueString: gofakeit.NamePrefix(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyUsualPrefix(listTesting["UsualPrefix"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetUsualPrefix(), listTesting["UsualPrefix"].ValueString)
		},
	}

	// ---- UsualAsPath ----
	listTesting["UsualAsPath"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ProductUseCase(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyUsualAsPath(listTesting["UsualAsPath"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetUsualAsPath(), listTesting["UsualAsPath"].ValueString)
		},
	}

	// ---- Source ----
	listTesting["Source"] = datamodelstest.TestOptions{
		ValueString: gofakeit.State(),
		SetFunc: func() {
			biZoneIRPCase.SetAnySource(listTesting["Source"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSource(), listTesting["Source"].ValueString)
		},
	}
	// ---- ExternalId ----
	listTesting["ExternalId"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyExternalId(listTesting["ExternalId"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetExternalId(), listTesting["ExternalId"].ValueString)
		},
	}

	// ---- UnderliningSource ----
	listTesting["UnderliningSource"] = datamodelstest.TestOptions{
		ValueString: gofakeit.EmojiAlias(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyUnderliningSource(listTesting["UnderliningSource"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetUnderliningSource(), listTesting["UnderliningSource"].ValueString)
		},
	}

	// ---- UpdatedAll ----
	listTesting["UpdatedAll"] = datamodelstest.TestOptions{
		ValueTime: gofakeit.Date(),
		SetFunc: func() {
			assert.NoError(t, biZoneIRPCase.SetAnyUpdatedAll(listTesting["UpdatedAll"].ValueTime.String()))
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetUpdatedAll(), listTesting["UpdatedAll"].ValueTime.Format(time.RFC3339))
		},
	}

	// ---- GossopkaKey ----
	listTesting["GossopkaKey"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyGossopkaKey(listTesting["GossopkaKey"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetGossopkaKey(), listTesting["GossopkaKey"].ValueString)
		},
	}

	// ---- ID ----
	listTesting["ID"] = datamodelstest.TestOptions{
		ValueUint64: gofakeit.Uint64(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyID(listTesting["ID"].ValueUint64)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetID(), listTesting["ID"].ValueUint64)
		},
	}

	// ---- System ----
	listTesting["System"] = datamodelstest.TestOptions{
		ValueUint64: gofakeit.Uint64(),
		SetFunc: func() {
			biZoneIRPCase.SetAnySystem(listTesting["System"].ValueUint64)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSystem(), listTesting["System"].ValueUint64)
		},
	}

	// ---- CancelGossopkaSendButton ----
	listTesting["CancelGossopkaSendButton"] = datamodelstest.TestOptions{
		ValueBool: gofakeit.Bool(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyCancelGossopkaSendButton(listTesting["CancelGossopkaSendButton"].ValueBool)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetCancelGossopkaSendButton(), listTesting["CancelGossopkaSendButton"].ValueBool)
		},
	}

	// ---- IsVisibleForCustomer ----
	listTesting["IsVisibleForCustomer"] = datamodelstest.TestOptions{
		ValueBool: gofakeit.Bool(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyIsVisibleForCustomer(listTesting["IsVisibleForCustomer"].ValueBool)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetIsVisibleForCustomer(), listTesting["IsVisibleForCustomer"].ValueBool)
		},
	}
	// ---- ShowGossopkaButton ----
	listTesting["ShowGossopkaButton"] = datamodelstest.TestOptions{
		ValueBool: gofakeit.Bool(),
		SetFunc: func() {
			biZoneIRPCase.SetShowGossopkaButton(listTesting["ShowGossopkaButton"].ValueBool)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetShowGossopkaButton(), listTesting["ShowGossopkaButton"].ValueBool)
		},
	}
	// ---- ShowGtiButton ----
	listTesting["ShowGtiButton"] = datamodelstest.TestOptions{
		ValueBool: gofakeit.Bool(),
		SetFunc: func() {
			biZoneIRPCase.SetAnyShowGtiButton(listTesting["ShowGtiButton"].ValueBool)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetShowGtiButton(), listTesting["ShowGtiButton"].ValueBool)
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

	/*
		t.Run("Test Any", func(t *testing.T) {
			timeStr := gofakeit.Date().Format(time.RFC3339)
			tt, err := time.Parse(time.RFC3339, timeStr) //"2026-02-19T16:33:03.108913+03:00")
			assert.NoError(t, err)
			assert.Equal(t, timeStr, tt.Format(time.RFC3339))

			fmt.Printf("Any time: '%+v', RFC3339 string: '%s'\n", tt, tt.Format(time.RFC3339))

			timeStr = "2026-02-19T16:33:03.108913+03:00"
			tt, err = time.Parse(time.RFC3339Nano, timeStr)
			assert.NoError(t, err)
			assert.Equal(t, timeStr, tt.Format(time.RFC3339))

		})
	*/
}

type DataStorage[T comparable] struct {
	Data T
}

func (storage *DataStorage[T]) Compare(a []DataStorage[T]) bool {
	for _, v := range a {
		if v.Data == storage.Data {
			return true
		}
	}

	return false
}
