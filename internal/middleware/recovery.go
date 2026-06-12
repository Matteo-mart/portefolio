package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

// Recovery est un middleware qui récupère les panics
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log le panic avec la stack trace
				log.Error().
					Interface("error", err).
					Str("path", r.URL.Path).
					Str("method", r.Method).
					Bytes("stack", debug.Stack()).
					Msg("Panic récupéré")

				// Renvoyer une erreur 500 propre
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"success": false, "error": "Une erreur interne s'est produite"}`)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
