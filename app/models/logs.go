package models

import (
	"time"

	"gorm.io/gorm"
)

type Logs struct {
	gorm.Model
	IP_Address   string    `gorm:"column:ip_address;size:255;"`
	Content      string    `gorm:"column:Content;size:255;"`
	Url_Callback string    `gorm:"column:url_callback;size:255;"`
	LoginAt      time.Time `gorm:"column:login_at;"`
}

func (r *Logs) Tablename() string {
	return "line_auth_logs"
}
