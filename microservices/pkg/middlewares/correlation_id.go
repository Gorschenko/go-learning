package middlewares

import (
	"context"
	"net/http"
	"pkg/logger"
	"pkg/static"

	"github.com/google/uuid"
)

func CorrelationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logger.GetLogger(ctx)

		correlationId := r.Header.Get(static.HeadersCorrelationID)

		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		logger.Debug("CorrelationID is created", "ID", correlationId)

		w.Header().Set(static.HeadersCorrelationID, correlationId)
		ctx = context.WithValue(ctx, static.ContextCorrelationID, correlationId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
