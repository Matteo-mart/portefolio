package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

// responseWriter wraps http.ResponseWriter pour capturer le status code
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

// Logger est un middleware qui log toutes les requêtes HTTP
func Logger(logger zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrapper pour capturer le status
			wrapped := &responseWriter{
				ResponseWriter: w,
				status:         200,
			}

			// Exécuter le handler
			next.ServeHTTP(wrapped, r)

			// Log après l'exécution
			duration := time.Since(start)

			logEvent := logger.Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Int("status", wrapped.status).
				Dur("duration", duration).
				Int("size", wrapped.size).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent())

			if r.URL.RawQuery != "" {
				logEvent.Str("query", r.URL.RawQuery)
			}

			logEvent.Msg("HTTP request")
		})
	}
}
