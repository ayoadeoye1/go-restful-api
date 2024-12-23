package models

import (
	"time"
)

type Posts struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Image   string `gorm:"type:varchar(512)" json:"image"`
	Tags    string `gorm:"type:varchar(255)" json:"tags"`
	Views   int    `gorm:"type:int;default:0" json:"views"`
	Likes   int    `gorm:"type:int;default:0" json:"likes"`
	UserID  uint   `gorm:"not null" json:"user_id"`

	// Time Stamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User *Users `gorm:"foreignKey:UserID" json:"user"`
}
