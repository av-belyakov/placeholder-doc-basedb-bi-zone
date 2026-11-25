package databasestorageapi

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// Start инициализирует новый модуль взаимодействия с API Database
func (dbs *DatabaseStorage) Start(ctx context.Context, reg *prometheus.Registry) error {
	baseTransport := &http.Transport{
		MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
		IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
		MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
		ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
		DialContext: (&net.Dialer{
			Timeout: 3 * time.Second,
			//KeepAlive: 1 * time.Second,
		}).DialContext,
	}

	es, err := elasticsearch.New(
		elasticsearch.WithAddresses(fmt.Sprintf("http://%s:%d", dbs.settings.host, dbs.settings.port)),
		elasticsearch.WithBasicAuth(dbs.settings.user, dbs.settings.passwd),
		elasticsearch.WithRetry(5, 429, 502, 503, 504), // не более 5 попыток повторного соединения если были получены перечисленые ответы
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(
				TransportWrapper(baseTransport, reg),
			),
		),
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(
				TransportWrapper(baseTransport, reg),
			),
		),
	)
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

// GetClientDB нативный клиент Elasticsearch
func (dbs *DatabaseStorage) GetClientDB() *elasticsearch.Client {
	return dbs.client
}

// Ping проверка соединения
func (dbs *DatabaseStorage) Ping() error {
	res, err := dbs.client.Ping()
	if res.StatusCode != http.StatusOK || err != nil {
		return fmt.Errorf("ping failed (status code is %d, error message-%s)", res.StatusCode, err.Error())
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

// RoundTrip реализует интерфейс http.RoundTripper
func (t *esTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var (
		code  string    = "error" // сетевая ошибка: таймаут, refused connection, TLS и т.д.
		start time.Time = time.Now()
	)

	t.metrics.inFlight.Inc() // Активных запросов к ES сейчас
	defer t.metrics.inFlight.Dec()
	resp, err := t.base.RoundTrip(req)

	t.metrics.inFlight.Dec()
	elapsed := time.Since(start).Seconds()
	if err == nil {
		code = strconv.Itoa(resp.StatusCode)
	}

	path := normalizePath(req.URL.Path)
	t.metrics.requests.WithLabelValues(req.Method, path, code).Inc()           // HTTP-запросы к Elasticsearch по методу, нормализованному пути и коду ответа
	t.metrics.queryDuration.WithLabelValues(req.Method, path).Observe(elapsed) // Длительность запросов к Elasticsearch.

	return resp, err
}
