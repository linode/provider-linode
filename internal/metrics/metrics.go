package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	k8smetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	metricsLinodeApiResponseCodesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "linode_api_responses_total",
		Help: "Number of Linode API responses by return code and first 5 digits of the token",
	}, []string{"service", "method", "code", "account"})

	metricsLinodeApiResponseCodesLast5m = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "linode_api_responses_last_5m",
		Help: "Number of Linode API responses by return code and first 5 digits of the token",
	}, []string{"service", "method", "code", "account"})
)

// SetupMetrics will register the known Prometheus metrics with controller-runtime's metrics registry
func SetupMetrics() error {
	k8smetrics.Registry.MustRegister(
		metricsLinodeApiResponseCodesTotal,
		metricsLinodeApiResponseCodesLast5m,
	)

	go func() {
		// Reset the counters every 5 minutes
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			metricsLinodeApiResponseCodesLast5m.Reset()
		}
	}()

	return nil
}

// IncLinodeAPIResp will increment the linode_api_responses_total metric for the specified service, operation, and responseCode tuple
func IncLinodeAPIResp(service, method, code, account string) error {
	metricsLinodeApiResponseCodesLast5m.WithLabelValues(service, method, code, account).Inc()
	metricsLinodeApiResponseCodesTotal.WithLabelValues(service, method, code, account).Inc()
	return nil
}
