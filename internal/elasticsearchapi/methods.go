package elasticsearchapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// Write запись сообщений
func (edb *ElasticsearchDB) Write(msgType, msg string) error {
	if edb.client == nil {
		return errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	msg = supportingfunctions.ReplaceCommaCharacter(msg)

	tn := time.Now()
	buf := bytes.NewReader(fmt.Appendf(nil, `{
		  "datetime": "%s",
		  "type": "%s",
		  "nameRegionalObject": "%s",
		  "message": "%s"
		}`,
		tn.Format(time.RFC3339),
		msgType,
		edb.settings.NameRegionalObject,
		msg,
	))

	res, err := edb.client.Index(fmt.Sprintf("logs.%s_%s_%d", edb.settings.IndexDB, strings.ToLower(tn.Month().String()), tn.Year()), buf)
	if err != nil {
		return supportingfunctions.CustomError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return nil
	}

	r := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return supportingfunctions.CustomError(err)
	}

	if e, ok := r["error"]; ok {
		return supportingfunctions.CustomError(fmt.Errorf("%s received from module Elsaticsearch: %s", res.Status(), e))
	}

	return nil
}
