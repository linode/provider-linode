package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	k8smetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	metricsLinodeApiCalls = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "linode_api_calls_total",
		Help: "Number of API calls to the Linode API",
	}, []string{"url", "method"})
)

// SetupMetrics will register the known Prometheus metrics with controller-runtime's metrics registry
func SetupMetrics() error {
	return k8smetrics.Registry.Register(metricsLinodeApiCalls)
}

// IncLinodeAPICall will increment the linode_api_calls_total metric for the specified service, operation, and apiVersion tuple
func IncLinodeAPICall(url, method string) {
	metricsLinodeApiCalls.WithLabelValues(url, method).Inc()
}
