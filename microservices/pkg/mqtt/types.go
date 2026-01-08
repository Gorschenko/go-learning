package mqtt

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	client  mqtt.Client
	options *mqtt.ClientOptions
}

type HandlerFunc func(ctx context.Context, message mqtt.Message)
