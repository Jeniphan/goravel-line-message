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
	err = db.AutoMigrate(
		&models.Users{},
		&models.UserAdmins{},
		&models.MessageTypes{},
		&models.UserMessageTypes{},
		&models.MessagesReceivedText{},
		&models.MessagesReceivedSticker{},
		&models.MessagesReceivedLocation{},
		&models.MessagesReceivedImage{},
		&models.MessagesReceivedAudio{},
		&models.RepliesMessage{},
		&models.Logs{},
	)
	if err != nil {
		panic(err)
	} else {
		createInitTypeMessage()
		createInitUserAdmin()
	}

	DB = db
}

func createInitTypeMessage() {
	messageTypes := []models.MessageTypes{
		{
			MessageTypeName: "Text",
			MessageTypeSlug: "text"},
		{
			MessageTypeName: "Sticker",
			MessageTypeSlug: "stricker"},
		{
			MessageTypeName: "Image",
			MessageTypeSlug: "img",
		},
		{
			MessageTypeName: "Audio",
			MessageTypeSlug: "audio",
		},
		{
			MessageTypeName: "Locations",
			MessageTypeSlug: "locations",
		},
	}

	for _, v := range messageTypes {
		err := facades.Orm().Query().UpdateOrCreate(&v, models.MessageTypes{MessageTypeSlug: v.MessageTypeSlug}, &v)
		if err != nil {
			panic(err)
		}
	}

	// err := facades.Orm().Query().Create(&messageTypes)
	// if err != nil {
	// 	panic(err)
	// }
}

func createInitUserAdmin() {
	userAdmin := models.UserAdmins{}
	userAdmin.UserName = "systems"
	userAdmin.FirstName = "Systems"
	userAdmin.LastName = "Axnos Line"
	userAdmin.Email = "systems.line-message@axonstech.com"

	err := facades.Orm().Query().UpdateOrCreate(&userAdmin, models.UserAdmins{UserName: "systems"}, &userAdmin)
	if err != nil {
		panic(err)
	}
}
