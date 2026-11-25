package databasestorageapi

import (
	"errors"
	"net/http"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/prometheus/client_golang/prometheus"
)

// New настраивает новый модуль взаимодействия с API Database
func New(logger interfaces.Logger, opts ...DatabaseStorageOptions) (*DatabaseStorage, error) {
	dbs := &DatabaseStorage{
		logger:   logger,
		chInput:  make(chan SettingsChanInput),
		chOutput: make(chan SettingsChanOutput),
		settings: settingsDatabaseStorage{
			maxGetDocumentsSize: 10,
		},
	}

	for _, opt := range opts {
		if err := opt(dbs); err != nil {
			return dbs, err
		}
	}

	return dbs, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		dbs.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		dbs.settings.port = v

		return nil
	}
}

// WithUser имя пользователя для доступа к БД
func WithUser(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		dbs.settings.user = v

		return nil
	}
}

// WithPasswd пароль для доступа к БД
func WithPasswd(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		dbs.settings.passwd = v

		return nil
	}
}

// WithNameDB наименование БД
func WithNameDB(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.namedb = v

		return nil
	}
}

// WithStorage наименование коллекции или индекса БД
func WithStorage(v map[string]string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.storages = v

		return nil
	}
}

// WithMaxGetDocumentSize максимальное количество запрашиваемых документов (по умолчанию 10)
func WithMaxGetDocumentSize(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.maxGetDocumentsSize = v

		return nil
	}
}

func TransportWrapper(rt http.RoundTripper, reg *prometheus.Registry) http.RoundTripper {
	esT := &esTransport{
		base: rt,
		metrics: metrics{
			inFlight: prometheus.NewGauge(
				prometheus.GaugeOpts{
					Name: "es_requests_in_flight",
					Help: "Активных запросов к ES сейчас.",
				},
			),
			requests: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "es_client_requests_total",
					Help: "HTTP-запросы к Elasticsearch по методу, нормализованному пути и коду ответа",
				},
				[]string{"method", "path", "code"},
			),
			queryDuration: prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name: "es_query_duration_seconds",
					Help: "Длительность запросов к Elasticsearch.",
					// поисковые запросы к ES обычно дольше SQL:
					// от ~1ms до единиц секунд
					Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
				},
				[]string{"operation", "index", "status"},
			),
			//hitsTotal: prometheus.NewCounterVec(
			//	prometheus.CounterOpts{
			//		Name: "es_query_hits_total",
			//		Help: "Суммарное число найденных документов.",
			//	},
			//	[]string{"index"},
			//),
			//emptyResults: prometheus.NewCounterVec(
			//	prometheus.CounterOpts{
			//		Name: "es_query_empty_results_total",
			//		Help: "Запросов с нулём результатов.",
			//	},
			//	[]string{"index"},
			//),
			//documentsIndexed: prometheus.NewCounterVec(
			//	prometheus.CounterOpts{
			//		Name: "es_documents_indexed_total",
			//		Help: "Проиндексированных документов.",
			//	},
			//	[]string{"index", "result"}, // created / updated / failed
			//),
		},
	}

	reg.MustRegister(
		esT.metrics.inFlight,
		esT.metrics.requests,
		esT.metrics.queryDuration,
		//esT.metrics.hitsTotal,
		//esT.metrics.emptyResults,
		//esT.metrics.documentsIndexed,
	)

	return esT
}
