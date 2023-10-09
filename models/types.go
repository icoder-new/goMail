package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	Firstname        string         `json:"firstname" gorm:"not null"`
	Lastname         string         `json:"lastname" gorm:"not null"`
	Email            string         `json:"email" gorm:"not null;unique"`
	MessageSentCount int            `json:"message_sent_count" gorm:"not null"`
	MessageSentDate  time.Time      `json:"message_sent" gorm:"not null"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

}
