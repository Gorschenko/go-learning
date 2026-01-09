package mqtt

import (
	"context"
	"encoding/json"
	"pkg/configs"
	"pkg/logger"
	"pkg/static"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func NewMqttService(config *configs.Config) (*MqttService, error) {
	ctx := context.Background()
	host := config.MQTT.Host
	port := config.MQTT.Port
	broker := "mqtt://" + host + ":" + strconv.Itoa(port)

	logger := logger.GetLogger(ctx)

	options := mqtt.
		NewClientOptions().
		AddBroker(broker).
		SetOnConnectHandler(func(client mqtt.Client) {
			logger.Info(
				"MqttService",
				"Connected to MQTT broker", broker,
			)
		}).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			logger.Info(
				"MqttService",
				"MQTT connection lost", err,
			)
		})

	client := mqtt.NewClient(options)

	service := MqttService{
		client,
		options,
	}

	if err := service.connect(); err != nil {
		return nil, err
	}

	return &service, nil
}

func (m *MqttService) connect() error {
	token := m.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *MqttService) Connect() error {
	return m.Connect()
}

func (m *MqttService) Disconnect() {
	m.client.Disconnect(250)
}

func (m *MqttService) Subscribe(topic string, qos byte, handler HandlerFunc) error {
	token := m.client.Subscribe(topic, qos, func(client mqtt.Client, message mqtt.Message) {
		ctx := context.Background()
		correlationId := uuid.New().String()
		ctx = context.WithValue(ctx, static.ContextCorrelationID, correlationId)
		logger := logger.GetLogger(ctx)

		logger.Info(
			"MQTTService",
			"Received message on topic", message.Topic(),
			"Messsage", message.Payload(),
		)

		handler(ctx, message)
	})

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *MqttService) Publish(ctx context.Context, topic string, qos byte, payload any) error {
	payloadToString, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	token := m.client.Publish(topic, qos, false, payloadToString)

	logger := logger.GetLogger(ctx)
	logger.Info(
		"MQTTService",
		"Published message on topic", topic,
		"Messsage", payloadToString,
	)

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
