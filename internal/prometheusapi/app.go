// Модуль для взаимодействия с API Prome
package prometheusapi

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// New настраивает модуль взаимодействия с API Prometheus
func New() (*PrometheusApi, error) {
	api := &PrometheusApi{
		settings: ApiSettings{
			connectionTimeout: 10,
		},
	}

	api.registry = prometheus.NewRegistry()
	api.registry.MustRegister(
		collectors.NewGoCollector(
			collectors.WithGoCollectorRuntimeMetrics(
				collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/sched/latencies:seconds")},
			),
		),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return api, nil
}
