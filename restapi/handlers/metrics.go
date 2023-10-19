package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/bnb-chain/mind-marketplace-backend/util"
)

var (
	namespace = "marketplace"

	counter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "endpoint_request_count",
		Help:      "Request count.",
	}, []string{"app", "name", "method", "state"})

	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "endpoint_duration_seconds",
		Help:      "Time taken to execute endpoint.",
	}, []string{"app", "name", "method", "status"})
)

func handleMetrics(handler http.Handler, app string) http.Handler {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)
		handler.ServeHTTP(rw, r)

		statusCode := rw.statusCode
		duration := time.Since(start)
		histogram.WithLabelValues(app, r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode)).Observe(duration.Seconds())
		counter.WithLabelValues(app, r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode)).Inc()
	})

	if err := prometheus.Register(histogram); err != nil {
		util.Logger.Errorf("prometheus.Register histogram error, err=%s", err.Error())
	}
	if err := prometheus.Register(counter); err != nil {
		util.Logger.Errorf("prometheus.Register counter error, err=%s", err.Error())
	}

	return h
}
