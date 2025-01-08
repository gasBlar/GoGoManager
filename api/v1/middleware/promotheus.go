package middleware

import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Total number of HTTP requests made.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_error_count",
			Help: "Total number of HTTP errors.",
		},
		[]string{"path", "status"},
	)
)

func PromotheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *CustomResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *CustomResponseWriter) Status() int {
	return rw.statusCode
}

func TrackMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := &CustomResponseWriter{ResponseWriter: w}
		path := r.URL.Path
		next.ServeHTTP(crw, r)
		status := crw.statusCode
		slog.Info("tracks metrics: ", "method", r.Method, "path", path, "status", http.StatusText(status))
		RequestCount.WithLabelValues(path, string(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, string(status)).Inc()
		}
	})
}
