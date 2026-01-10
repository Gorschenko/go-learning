package main

import (
	"context"
	"devices/internal/devices"
	"math/rand"
	"pkg/configs"
	"pkg/logger"
	"pkg/mqtt"
	mqtt_devices_api "pkg/mqtt/devices"
	"pkg/static"
	"strconv"
	"time"

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

		serialNumber := strconv.Itoa(rand.Intn(1000))
		update := mqtt_devices_api.DeviceUpdateDto{
			Status: devices.GetRandomDeviceStatus(),
		}

		mqttDevicesApi.SendDeviceUpdateEvent(ctx, serialNumber, &update)

		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}
