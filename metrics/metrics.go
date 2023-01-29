package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func StartMetricsServer(metricsServerAddress string) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(metricsServerAddress, nil)
}
