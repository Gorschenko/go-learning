package mqtt_middlewares

import (
	pkg_mqtt "pkg/mqtt"
)

func CombineMiddlewares(middlewares ...pkg_mqtt.Middleware) pkg_mqtt.Middleware {
	return func(next pkg_mqtt.Handler) pkg_mqtt.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next
	}
}
