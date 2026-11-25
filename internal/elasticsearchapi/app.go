package elasticsearchapi

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

// New соединение с базой данных
func New(settings Settings) (*ElasticsearchDB, error) {
	edb := &ElasticsearchDB{
		settings: settings,
	}

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
		elasticsearch.WithAddresses(fmt.Sprintf("http://%s:%d", settings.Host, settings.Port)),
		elasticsearch.WithBasicAuth(settings.User, settings.Passwd),
		elasticsearch.WithRetry(5, 429, 502, 503, 504), // не более 5 попыток повторного соединения если были получены перечисленые ответы
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(baseTransport),
		),
	)
	if err != nil {
		return edb, err
	}
	edb.client = es

	return edb, err
}
