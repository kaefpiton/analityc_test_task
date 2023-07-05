package http

import (
	"analityc_test_task/internal/metrics"
	"analityc_test_task/pkg/logger"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func startHTTPMetricServer(port string, log logger.Logger) {
	metricServer := echo.New()
	metricGroup := metricServer.Group(
		"/metrics",
	)

	err := metrics.RegisterMetrics()
	if err != nil {
		log.ErrorF("Failed to register metrics: %v", err.Error())
	}

	h := promhttp.Handler()
	metricGroup.GET("", func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	go func() {
		port := fmt.Sprintf(":%v", port)
		if err = metricServer.Start(port); err != nil {
			log.Error(err)
		}
	}()
}
