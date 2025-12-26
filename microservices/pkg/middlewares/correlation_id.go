package middlewares

import (
	"context"
	"net/http"
	"pkg/static"

	"github.com/google/uuid"
)

func CorrelationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := r.Header.Get(static.HeadersCorrelationID)

		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), static.ContextCorrelationID, correlationId)
		w.Header().Set(static.HeadersCorrelationID, correlationId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
