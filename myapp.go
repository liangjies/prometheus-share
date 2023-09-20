package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

// NewGauge

func recordMetrics() {
	go func() {
		opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "The total number of processed events",
		})
		requestCount := promauto.NewGauge(prometheus.GaugeOpts{
			Name: "http_request_duration_seconds",
			Help: "Histogram of lantencies for HTTP requests",
		})
		//goInfo := prometheus.NewGaugeVec(
		//	prometheus.GaugeOpts{
		//		Name: "go_info_2",
		//		Help: "Go version information",
		//	},
		//	[]string{"version"},
		//)
		//prometheus.MustRegister(goInfo)
		//goInfo.With(prometheus.Labels{"version": "go1.20.8"}).Set(1)
		for {
			opsProcessed.Inc()
			requestCount.Dec()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
