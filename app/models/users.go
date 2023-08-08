package models

import (
	"github.com/golang-module/carbon/v2"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type Users struct {
	orm.Model
	Id                       uint                        `gorm:"primaryKey" json:"id"`
	UserLineID               string                      `gorm:"size:255;column:user_line_id" form:"user_line_id" json:"user_line_id"`
	UserType                 string                      `gorm:"size:255;column:user_type" form:"user_type" json:"user_type"`
	CreateAt                 carbon.DateTime             `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt                carbon.DateTime             `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt                gorm.DeletedAt              `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	MessagesReceivedText     []*MessagesReceivedText     `gorm:"foreignKey:UserID"`
	MessagesReceivedSticker  []*MessagesReceivedSticker  `gorm:"foreignKey:UserID"`
	MessagesReceivedLocation []*MessagesReceivedLocation `gorm:"foreignKey:UserID"`
	MessagesReceivedImage    []*MessagesReceivedImage    `gorm:"foreignKey:UserID"`
	MessagesReceivedAudio    []*MessagesReceivedAudio    `gorm:"foreignKey:UserID"`
}

type MessagesReceivedText struct {
	orm.Model
	Id uint `gorm:"primaryKey" json:"id"`

	MessageText   string          `gorm:"size:255;message_text;column:message_text" json:"message_text"`
	UserID        uint            `gorm:"column:user_id" json:"user_id"`
	MessageLineID string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt      carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt     carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users         *Users          `gorm:"foreignKey:UserID"`
}

type MessagesReceivedSticker struct {
	orm.Model
	Id                  uint            `gorm:"primaryKey" json:"id"`
	StickerId           string          `gorm:"size:255;sticker_id;column:sticker_id" json:"sticker_id"`
	PackageId           string          `gorm:"size:255;package_id;column:package_id" json:"package_id"`
	StickerResourceType string          `gorm:"size:255;sticker_resource_type;column:sticker_resource_type" json:"sticker_resource_type"`
	Keywords            string          `gorm:"size:255;key_words;column:key_words" json:"key_words"`
	UserID              uint            `gorm:"column:user_id" json:"user_id"`
	MessageLineID       string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt            carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt           carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users               *Users          `gorm:"foreignKey:UserID"`
}

type MessagesReceivedLocation struct {
	orm.Model
	Id            uint            `gorm:"primaryKey" json:"id"`
	Address       string          `gorm:"size:255;address;column:address" json:"address"`
	Latitude      string          `gorm:"size:255;latitude;column:latitude" json:"latitude"`
	Longitude     string          `gorm:"size:255;longitude;column:longitude" json:"longitude"`
	UserID        uint            `gorm:"column:user_id" json:"user_id"`
	MessageLineID string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt      carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt     carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users         *Users          `gorm:"foreignKey:UserID"`
}

type MessagesReceivedImage struct {
	orm.Model
	Id              uint            `gorm:"primaryKey" json:"id"`
	ContentProvider string          `gorm:"size:255;content_provider;column:content_provider" json:"content_provider"`
	UserID          uint            `gorm:"column:user_id" json:"user_id"`
	MessageLineID   string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt        carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt       carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users           *Users          `gorm:"foreignKey:UserID"`
}

type MessagesReceivedAudio struct {
	orm.Model
	Id              uint            `gorm:"primaryKey" json:"id"`
	ContentProvider string          `gorm:"size:255;content_provider;column:content_provider" json:"content_provider"`
	Duration        string          `gorm:"size:255;duration;column:duration" json:"duration"`
	UserID          uint            `gorm:"column:user_id" json:"user_id"`
	MessageLineID   string          `gorm:"size:255;message_line_id;column:message_line_id" json:"message_line_id"`
	CreateAt        carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt       carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Users           *Users          `gorm:"foreignKey:UserID"`
}

type RepliesMessage struct {
	orm.Model
	Id                         uint                      `gorm:"primaryKey" json:"id"`
	MessageText                string                    `gorm:"size:255;message_text;column:message_text" json:"message_text"`
	UserID                     uint                      `gorm:"column:user_id" json:"user_id"`
	MessagesReceivedTextID     *uint                     `gorm:"column:messages_received_text_id;default:NULL" json:"messages_received_text_id"`
	MessagesReceivedStickerID  *uint                     `gorm:"column:messages_received_sticker_id;default:NULL" json:"messages_received_sticker_id"`
	MessagesReceivedLocationID *uint                     `gorm:"column:messages_received_location_id;default:NULL" json:"messages_received_location_id"`
	MessagesReceivedImageID    *uint                     `gorm:"column:messages_received_image_id;default:NULL" json:"messages_received_image_id"`
	MessagesReceivedAudioID    *uint                     `gorm:"column:messages_received_audio_id;default:NULL" json:"messages_received_audio_id"`
	CreateAt                   carbon.DateTime           `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt                  carbon.DateTime           `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt                  gorm.DeletedAt            `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	MessagesReceivedText       *MessagesReceivedText     `gorm:"foreignKey:MessagesReceivedTextID"`
	MessagesReceivedSticker    *MessagesReceivedSticker  `gorm:"foreignKey:MessagesReceivedStickerID"`
	MessagesReceivedLocation   *MessagesReceivedLocation `gorm:"foreignKey:MessagesReceivedLocationID"`
	MessagesReceivedImage      *MessagesReceivedImage    `gorm:"foreignKey:MessagesReceivedImageID"`
	MessagesReceivedAudio      *MessagesReceivedAudio    `gorm:"foreignKey:MessagesReceivedAudioID"`
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
