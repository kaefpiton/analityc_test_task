package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "REQUESTS_TOTAL",
	})
)

func RegisterMetrics() error {
	var collectors = []prometheus.Collector{
		RequestsTotal,
	}

	for i := range collectors {
		if err := prometheus.Register(collectors[i]); err != nil {
			return fmt.Errorf("error register metric (index=%d). %w", i, err)
		}

	}
	return nil
}
