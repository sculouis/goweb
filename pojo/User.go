package pojo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	PassWord string `json:"UserPassWord"`
	Email    string `json:"UserEmail"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
}

var DB *gorm.DB

func ConnectDatabase() {
	var dsn = "host=postgres user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})

	DB = database
}
