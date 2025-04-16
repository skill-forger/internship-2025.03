package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel represents the fundamental fields in all database tables
type BaseModel struct {
	ID        int `gorm:"TYPE:BIGINT(11);UNSIGNED;AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
