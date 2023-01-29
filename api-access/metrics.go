package api_access

import "github.com/prometheus/client_golang/prometheus"

var (
	apiAccessQPS *prometheus.CounterVec = nil
	apiAccessErr *prometheus.CounterVec = nil
)

func InitMetrics() {
	apiAccessQPS = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_access_qps",
	}, []string{"service", "method"})
	apiAccessErr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_access_err",
	}, []string{"service", "method", "errmsg"})

	prometheus.MustRegister(apiAccessQPS, apiAccessErr)
}
