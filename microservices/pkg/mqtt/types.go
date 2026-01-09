package mqtt

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	client  mqtt.Client
	options *mqtt.ClientOptions
}

type Middleware func(Handler) Handler

type Handler func(ctx context.Context, message mqtt.Message)
