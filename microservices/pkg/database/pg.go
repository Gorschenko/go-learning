package database

import (
	"pkg/configs"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) (*Db, error) {
	host := config.Database.Host
	port := config.Database.Port
	user := config.Database.User
	password := config.Database.Password
	database := config.Database.Database

	logLevel := logger.Silent

	dsn := "host=" + host + " " + "port=" + strconv.Itoa(port) + " " + "user=" + user + " " + "password=" + password + " " + "database=" + database + " " + "sslmode=disable"

	if config.Software.Logger.Level == "debug" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}
