package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt   gorm.DeletedAt
}
