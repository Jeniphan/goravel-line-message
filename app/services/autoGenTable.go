package services

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/facades"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func GenTable() {
	config := facades.Config()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable",
		config.GetString("database.connections.postgresql.host"),
		config.GetString("database.connections.postgresql.port"),
		config.GetString("database.connections.postgresql.username"),
		config.GetString("database.connections.postgresql.password"),
		config.GetString("database.connections.postgresql.database"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Users{},
	)
	db.AutoMigrate(
		&models.MessagesReceivedText{},
	)
	db.AutoMigrate(
		&models.MessagesReceivedSticker{},
	)
	db.AutoMigrate(
		&models.MessagesReceivedLocation{},
	)
	db.AutoMigrate(
		&models.MessagesReceivedImage{},
	)
	db.AutoMigrate(
		&models.MessagesReceivedAudio{},
	)
	db.AutoMigrate(
		&models.RepliesMessage{},
	)
	DB = db
}
