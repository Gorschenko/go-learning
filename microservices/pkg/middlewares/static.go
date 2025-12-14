package middlewares

type contextKey struct{}

var (
	ContextBodyKey   = contextKey{}
	ContextParamsKey = contextKey{}
)
