package model

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Level string `gorm:"type:varchar(255)"`
	Message string `gorm:"type:varchar(255)"`
	Timestamp time.Time `gorm:"type:timestamp"`
	Tag string `gorm:"type:varchar(255)"`
	Data string `gorm:"type:jsonb"`
}
