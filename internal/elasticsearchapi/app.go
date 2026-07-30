package elasticsearchapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// New соединение с базой данных
func New(settings Settings) (*ElasticsearchDB, error) {
	edb := &ElasticsearchDB{settings: settings}

	es, err := elasticsearch.New(
		elasticsearch.WithAddresses(fmt.Sprintf("http://%s:%d", settings.Host, settings.Port)),
		elasticsearch.WithBasicAuth(settings.User, settings.Passwd),
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(&http.Transport{
				MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
				IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
				MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
				ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
				DialContext: (&net.Dialer{
					Timeout: 3 * time.Second,
					//KeepAlive: 1 * time.Second,
				}).DialContext,
			},
			),
		),
	)
	if err != nil {
		return edb, err
	}

	edb.client = es

	return edb, err
}

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
