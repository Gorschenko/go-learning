package utils

import (
	"context"
	"log"
	"pkg/static"
)

func GetCorrelationID(ctx context.Context) string {
	ID, ok := ctx.Value(static.ContextCorrelationID).(string)

	if ok {
		return ID
	}
	return ""
}

func LogWithCorrelationId(class, message string) {
	correlationID := GetCorrelationID(context.Background())
	log.Println(correlationID, class, message)
}
