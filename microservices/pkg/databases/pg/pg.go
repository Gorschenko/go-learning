package databases

import (
	"pkg/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *config.Config) (*Db, error) {
	host := config.Databases.Pg.Host
	port := config.Databases.Pg.Port
	user := config.Databases.Pg.User
	password := config.Databases.Pg.Password
	database := config.Databases.Pg.Database

	dsn := "host=" + host + " " + "port=" + strconv.Itoa(port) + " " + "user=" + user + " " + "password=" + password + " " + "database=" + database + " " + "sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}
