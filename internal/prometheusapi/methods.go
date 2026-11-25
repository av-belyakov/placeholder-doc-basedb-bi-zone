package prometheusapi

import "github.com/prometheus/client_golang/prometheus"

// GetRegistry
func (p *PrometheusApi) GetRegistry() *prometheus.Registry {
	return p.registry
}
