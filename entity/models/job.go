package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	Quota       uint      `gorm:"not null"`
	Status      string    `gorm:"not null"`
	ExpiryDate  time.Time `gorm:"not null"`
	JobPosterID uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:JobPosterID;references:ID"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt   gorm.DeletedAt
}
