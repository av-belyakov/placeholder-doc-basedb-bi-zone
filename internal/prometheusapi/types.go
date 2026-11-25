package prometheusapi

import "github.com/prometheus/client_golang/prometheus"

type PrometheusApi struct {
	registry *prometheus.Registry
	settings ApiSettings
}

// settings настройки модуля
type ApiSettings struct {
	apiKey            string
	host              string
	port              int
	connectionTimeout int
}

// options функциональные опции
type options func(*PrometheusApi) error
