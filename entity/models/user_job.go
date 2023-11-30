package models

import (
	"time"

	"gorm.io/gorm"
)

type UserJob struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	JobID     uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	Job       Job       `gorm:"foreignKey:JobID;references:ID"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}
