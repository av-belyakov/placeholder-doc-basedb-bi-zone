package databasestorageapi

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// Start инициализирует новый модуль взаимодействия с API Database
func (dbs *DatabaseStorage) Start(ctx context.Context) error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", dbs.settings.host, dbs.settings.port)},
		Username:  dbs.settings.user,
		Password:  dbs.settings.passwd,
		Transport: &http.Transport{
			MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
			MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
			IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
			ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
				//KeepAlive: 1 * time.Second,
			}).DialContext,
		},
		//RetryOnError: ,
		//RetryOnStatus: ,
	})
	if err != nil {
		return supportingfunctions.CustomError(err)
	}

	dbs.client = es

	go dbs.router(ctx)

	return nil
}

// GetChannelToModule канал для передачи данных в модуль
func (dbs *DatabaseStorage) GetChannelToModule() chan SettingsChanInput {
	return dbs.chInput
}

// GetChannelFromModule канал для приёма данных из модуля
func (dbs *DatabaseStorage) GetChannelFromModule() chan SettingsChanOutput {
	return dbs.chOutput
}

// Ping проверка соединения
func (dbs *DatabaseStorage) Ping() error {
	res, err := dbs.client.Ping()
	if res.StatusCode != http.StatusOK || err != nil {
		return fmt.Errorf("ping failed (status code is %s, error message-%s)", res.StatusCode, err.Error())
	}

	return nil
}

// Update обёртка для обновления документа
func (dbs *DatabaseStorage) Update(index, underlineId string, bodyUpdate *strings.Reader) ([]byte, int, error) {
	esRes, err := dbs.client.Update(index, underlineId, bodyUpdate)
	if err != nil {
		return []byte{}, 500, err
	}
	defer esRes.Body.Close()

	res, err := io.ReadAll(esRes.Body)
	if err != nil {
		return []byte{}, esRes.StatusCode, err
	}

	return res, esRes.StatusCode, nil
}

// ConnectionClose закрытие соединения
func (dbs *DatabaseStorage) ConnectionClose() {}
