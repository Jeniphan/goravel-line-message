package models

import (
	"github.com/golang-module/carbon/v2"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type LineConfigs struct {
	gorm.Model
	Id                     uint            `gorm:"primaryKey" json:"id"`
	LineChannelId          string          `gorm:"size:255;column:line_channel_id" form:"line_channel_id" json:"line_channel_id"`
	LineChannelName        string          `gorm:"size:255;column:line_channel_name" form:"line_channel_name" json:"line_channel_name"`
	LineChannelSecret      string          `gorm:"size:255;column:line_channel_secret" form:"line_channel_secret" json:"line_channel_secret"`
	LineChannelAccessToken string          `gorm:"size:255;column:line_channel_access_token" form:"line_channel_access_token" json:"line_channel_access_token"`
	CreateAt               carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt              carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt              gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`

	MessagesReceivedText     []*MessagesReceivedText     `gorm:"foreignKey:LineConfigID"`
	MessagesReceivedSticker  []*MessagesReceivedSticker  `gorm:"foreignKey:LineConfigID"`
	MessagesReceivedLocation []*MessagesReceivedLocation `gorm:"foreignKey:LineConfigID"`
	MessagesReceivedImage    []*MessagesReceivedImage    `gorm:"foreignKey:LineConfigID"`
	MessagesReceivedAudio    []*MessagesReceivedAudio    `gorm:"foreignKey:LineConfigID"`
	UserMessageTypes         []*UserMessageTypes         `gorm:"foreignKey:LineConfigID"`
	Users                    []*Users                    `gorm:"foreignKey:LineConfigID"`
}

type Users struct {
	orm.Model
	Id                       uint                        `gorm:"primaryKey" json:"id"`
	UserLineID               string                      `gorm:"size:255;column:user_line_id" form:"user_line_id" json:"user_line_id"`
	UserType                 string                      `gorm:"size:255;column:user_type" form:"user_type" json:"user_type"`
	LineConfigID             uint                        `gorm:"column:line_config_id" json:"line_config_id"`
	CreateAt                 carbon.DateTime             `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt                carbon.DateTime             `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt                gorm.DeletedAt              `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	LineConfigs              *LineConfigs                `gorm:"foreignKey:LineConfigID"`
	MessagesReceivedText     []*MessagesReceivedText     `gorm:"foreignKey:UserID"`
	MessagesReceivedSticker  []*MessagesReceivedSticker  `gorm:"foreignKey:UserID"`
	MessagesReceivedLocation []*MessagesReceivedLocation `gorm:"foreignKey:UserID"`
	MessagesReceivedImage    []*MessagesReceivedImage    `gorm:"foreignKey:UserID"`
	MessagesReceivedAudio    []*MessagesReceivedAudio    `gorm:"foreignKey:UserID"`
	UserMessageTypes         []*UserMessageTypes         `gorm:"foreignKey:UserID"`
}

type MessagesReceivedText struct {
	orm.Model
	Id uint `gorm:"primaryKey" json:"id"`

	MessageText   string          `gorm:"size:255;message_text;column:message_text" json:"message_text"`
	UserID        uint            `gorm:"column:user_id" json:"user_id"`
	LineConfigID  uint            `gorm:"column:line_config_id" json:"line_config_id"`
	MessageLineID string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt      carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt     carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users         *Users          `gorm:"foreignKey:UserID"`
	LineConfigs   *LineConfigs    `gorm:"foreignKey:LineConfigID"`
}

type MessagesReceivedSticker struct {
	orm.Model
	Id                  uint            `gorm:"primaryKey" json:"id"`
	StickerId           string          `gorm:"size:255;sticker_id;column:sticker_id" json:"sticker_id"`
	PackageId           string          `gorm:"size:255;package_id;column:package_id" json:"package_id"`
	StickerResourceType string          `gorm:"size:255;sticker_resource_type;column:sticker_resource_type" json:"sticker_resource_type"`
	Keywords            string          `gorm:"size:255;key_words;column:key_words" json:"key_words"`
	UserID              uint            `gorm:"column:user_id" json:"user_id"`
	LineConfigID        uint            `gorm:"column:line_config_id" json:"line_config_id"`
	MessageLineID       string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt            carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt           carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users               *Users          `gorm:"foreignKey:UserID"`
	LineConfigs         *LineConfigs    `gorm:"foreignKey:LineConfigID"`
}

type MessagesReceivedLocation struct {
	orm.Model
	Id            uint            `gorm:"primaryKey" json:"id"`
	Address       string          `gorm:"size:255;address;column:address" json:"address"`
	Latitude      string          `gorm:"size:255;latitude;column:latitude" json:"latitude"`
	Longitude     string          `gorm:"size:255;longitude;column:longitude" json:"longitude"`
	UserID        uint            `gorm:"column:user_id" json:"user_id"`
	LineConfigID  uint            `gorm:"column:line_config_id" json:"line_config_id"`
	MessageLineID string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt      carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt     carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users         *Users          `gorm:"foreignKey:UserID"`
	LineConfigs   *LineConfigs    `gorm:"foreignKey:LineConfigID"`
}

type MessagesReceivedImage struct {
	orm.Model
	Id              uint            `gorm:"primaryKey" json:"id"`
	ContentProvider string          `gorm:"size:255;content_provider;column:content_provider" json:"content_provider"`
	UserID          uint            `gorm:"column:user_id" json:"user_id"`
	LineConfigID    uint            `gorm:"column:line_config_id" json:"line_config_id"`
	MessageLineID   string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt        carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt       carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users           *Users          `gorm:"foreignKey:UserID"`
	LineConfigs     *LineConfigs    `gorm:"foreignKey:LineConfigID"`
}

type MessagesReceivedAudio struct {
	orm.Model
	Id              uint   `gorm:"primaryKey" json:"id"`
	ContentProvider string `gorm:"size:255;content_provider;column:content_provider" json:"content_provider"`
	Duration        string `gorm:"size:255;duration;column:duration" json:"duration"`
	UserID          uint   `gorm:"column:user_id" json:"user_id"`
	LineConfigID    uint   `gorm:"column:line_config_id" json:"line_config_id"`

	MessageLineID string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt      carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt     carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users         *Users          `gorm:"foreignKey:UserID"`
	LineConfigs   *LineConfigs    `gorm:"foreignKey:LineConfigID"`
}

type RepliesMessage struct {
	orm.Model
	Id                 uint              `gorm:"primaryKey" json:"id"`
	MessageText        string            `gorm:"size:255;message_text;column:message_text" json:"message_text"`
	UserID             uint              `gorm:"column:user_id" json:"user_id"`
	LineConfigID       uint              `gorm:"column:line_config_id" json:"line_config_id"`
	CreateAt           carbon.DateTime   `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt          carbon.DateTime   `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt          gorm.DeletedAt    `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	UserAdminID        uint              `gorm:"column:user_admin_id" json:"user_admin_id"`
	UserMessageTypesID uint              `gorm:"column:user_message_type_id" json:"user_message_type_id"`
	UserMessageTypes   *UserMessageTypes `gorm:"foreignKey:UserMessageTypesID"`
	LineConfigs        *LineConfigs      `gorm:"foreignKey:LineConfigID"`
}

type UserMessageTypes struct {
	orm.Model
	Id             uint              `gorm:"primaryKey" json:"id"`
	UserID         uint              `gorm:"column:user_id" json:"user_id"`
	LineConfigID   uint              `gorm:"column:line_config_id" json:"line_config_id"`
	MessageId      uint              `gorm:"column:message_id" json:"message_id"`
	MessageTypeID  uint              `gorm:"column:message_type_id" json:"message_type_id"`
	CreateAt       carbon.DateTime   `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt      carbon.DateTime   `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt    `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users          *Users            `gorm:"foreignKey:UserID"`
	LineConfigs    *LineConfigs      `gorm:"foreignKey:LineConfigID"`
	MessageTypes   *MessageTypes     `gorm:"foreignKey:MessageTypeID"`
	RepliesMessage []*RepliesMessage `gorm:"foreignKey:UserMessageTypesID"`
}

type MessageTypes struct {
	orm.Model
	Id               uint                `gorm:"primaryKey" json:"id"`
	MessageTypeName  string              `gorm:"size:255;column:message_type_name" json:"message_type_name"`
	MessageTypeSlug  string              `gorm:"size:255;column:message_type_slug" json:"message_type_slug"`
	CreateAt         carbon.DateTime     `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt        carbon.DateTime     `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt      `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	UserMessageTypes []*UserMessageTypes `gorm:"foreignKey:MessageTypeID"`
}

type UserAdmins struct {
	orm.Model
	Id             uint              `gorm:"primaryKey" json:"id"`
	UserName       string            `gorm:"size:255;column:user_name" json:"user_name"`
	FirstName      string            `gorm:"size:255;column:first_name" json:"first_name"`
	LastName       string            `gorm:"size:255;column:last_name" json:"last_name"`
	Email          string            `gorm:"size:255;column:email" json:"email"`
	RepliesMessage []*RepliesMessage `gorm:"foreignKey:UserAdminID"`
	CreateAt       carbon.DateTime   `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt      carbon.DateTime   `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt    `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (m *MessageTypes) MessageType() string {
	return "message_types"
}

func (u *UserMessageTypes) UserMessageTypes() string {
	return "user_message_types"
}

func (r *Users) Users() string {
	return "users"
}

func (r *MessagesReceivedText) MessagesReceivedTexts() string {
	return "messages_received_text"
}

func (r *MessagesReceivedSticker) MessagesReceivedStickers() string {
	return "messages_received_sticker"
}

func (r *MessagesReceivedLocation) MessagesReceivedLocations() string {
	return "messages_received_location"
}

func (r *MessagesReceivedImage) MessagesReceivedImages() string {
	return "messages_received_image"
}

func (r *MessagesReceivedAudio) MessagesReceivedAudios() string {
	return "messages_received_audio"
}

func (r *RepliesMessage) RepliesMessages() string {
	return "replies_message"
}
