package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(data []byte) (int, error) {
	lrw.body.Write(data) // Capture response body
	return lrw.ResponseWriter.Write(data)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only log specific requests
		if r.URL.Path == "/v1/department" && r.Method == "PATCH" {
			// Read and log the request body
			var requestBody bytes.Buffer
			if r.Body != nil {
				// Copy and restore the request body
				_, err := io.Copy(&requestBody, r.Body)
				if err != nil {
					slog.Error("Error reading request body", "error", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				r.Body = io.NopCloser(&requestBody) // Restore the body for downstream handlers
			}

			slog.Info("Request received",
				"method", r.Method,
				"url", r.URL.Path,
				"body", requestBody.String(),
			)

			// Wrap the ResponseWriter
			lrw := &loggingResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK, // Default status
			}

			// Process the request
			next.ServeHTTP(lrw, r)

			// Log the response details
			slog.Info("Response sent",
				"status", lrw.statusCode,
				"body", lrw.body.String(),
				"headers", lrw.Header(),
			)
			return
		}

		// For all other requests, just proceed without logging
		next.ServeHTTP(w, r)
	})
}
