package static

type contextKey struct{}

var (
	ContextBodyKey       = contextKey{}
	ContextParamsKey     = contextKey{}
	ContextCorrelationID = contextKey{}
)
