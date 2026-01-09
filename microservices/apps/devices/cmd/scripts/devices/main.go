package main

import (
	"context"
	"math/rand"
	"pkg/configs"
	"pkg/database"
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
			Status: getRandomDeviceStatus(),
		}

		mqttDevicesApi.SendUpdateDeviceEvent(ctx, serialNumber, &update)

		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func getRandomDeviceStatus() database.DeviceStatus {
	num := rand.Intn(100)

	if num <= 10 {
		return database.DevicesStatusUnknown
	} else if num > 10 && num <= 50 {
		return database.DeviceStatusOnline
	} else {
		return database.DeviceStatusOffline
	}
}
