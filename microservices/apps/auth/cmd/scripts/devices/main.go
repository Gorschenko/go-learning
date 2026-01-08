package main

import (
	"context"
	"math/rand"
	"pkg/configs"
	"pkg/logger"
	"pkg/mqtt"
	mqtt_devices_api "pkg/mqtt/devices"
	"pkg/static"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func main() {
	config, err := configs.LoadConfig("../../config.json")
	if err != nil {
		panic(err)
	}

	logger.SetupLogger(config)

	mqttService, err := mqtt.NewMqttService(config)
	if err != nil {
		panic(err)
	}

	mqttDevicesApi := mqtt_devices_api.NewDevicesApi(&mqtt_devices_api.DevicesDependencies{
		MqttService: mqttService,
	})

	for {
		ctx := context.Background()
		correlationId := uuid.New().String()
		ctx = context.WithValue(ctx, static.ContextCorrelationID, correlationId)

		logger := logger.GetLogger(ctx)

		logger.Info("Create device event")

		serialNumber := strconv.Itoa(gofakeit.Int())

		update := mqtt_devices_api.DeviceUpdateDto{
			Status: getRandomDeviceStatus(),
		}
		mqttDevicesApi.SendUpdateDeviceEvent(ctx, serialNumber, &update)

		time.Sleep(2 * time.Second)
	}
}

func getRandomDeviceStatus() mqtt_devices_api.DeviceStatus {
	num := rand.Intn(100)

	if num <= 50 {
		return mqtt_devices_api.DeviceStatusOffline
	} else {
		return mqtt_devices_api.DeviceStatusOnline
	}
}
