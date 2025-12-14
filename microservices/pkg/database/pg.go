package database

import (
	"pkg/configs"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=" + host + " " + "port=" + strconv.Itoa(port) + " " + "user=" + user + " " + "password=" + password + " " + "database=" + database + " " + "sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}
