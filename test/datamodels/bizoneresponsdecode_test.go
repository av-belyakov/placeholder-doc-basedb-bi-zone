package datamodels_test

import (
	"encoding/json/v2"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
	"github.com/stretchr/testify/assert"
)

func TestBiZoneResponseDecode(t *testing.T) {
	data, err := os.ReadFile("../test_json/incident-responce.json")
	if err != nil {
		log.Fatalln(err)
	}

	var response databasestorageapi.ResponseVerifiedIRPIncidents
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatalln(err)
	}

	assert.NotEmpty(t, response.Options.Hits)
	assert.Greater(t, response.Options.Total.Value, 0)

	fmt.Println("Found size incident:", response.Options.Total.Value)

	//fmt.Printf("Result decode:'%+v'\n", response)
}
