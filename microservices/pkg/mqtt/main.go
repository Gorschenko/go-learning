package mqtt

import (
	"pkg/configs"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	client  mqtt.Client
	options *mqtt.ClientOptions
}

func NewMQTTService(config *configs.Config) *MQTTService {
	host := config.MQTT.Host
	port := config.MQTT.Port
	broker := "tcp://" + host + ":" + strconv.Itoa(port)

	options := mqtt.
		NewClientOptions().
		AddBroker(broker)

	client := mqtt.NewClient(options)

	return &MQTTService{
		client,
		options,
	}
}

func (m *MQTTService) Connect() error {
	token := m.client.Connect()
	token.Wait()
	return token.Error()
}
