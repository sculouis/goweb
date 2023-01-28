package db

import (
	. "GoWeb/pojo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dsn = "host=postgres user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(Todo{})
	database.AutoMigrate(User{})

	DB = database
}
