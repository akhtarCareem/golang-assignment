package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	logger := NewLogger()

	// Start Prometheus metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logger.Info("Prometheus metrics available at /metrics")
		http.ListenAndServe(":2112", nil)
	}()
}
