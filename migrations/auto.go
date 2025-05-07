package main

import (
	"go/adv-api/internal/link"
	"go/adv-api/internal/stat"
	"go/adv-api/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().AutoMigrate(&link.Link{}, &user.User{},&stat.Stat{})

}
