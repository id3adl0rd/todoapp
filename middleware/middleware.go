package middleware

import (
	"net/http"
	"runtime/debug"
	"to-do-app/logger"
	"to-do-app/route"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				route.RespondWithError(w, http.StatusInternalServerError, "occurred error")
				logger.Log.Errorf("Recovered from panic: %s", string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
