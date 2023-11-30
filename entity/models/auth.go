package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null;unique"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}
