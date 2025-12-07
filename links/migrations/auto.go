package main

import (
	"os"
	"test/internal/links"
	"test/internal/stats"
	"test/internal/users"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&links.Link{})
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&stats.Stat{})
}
