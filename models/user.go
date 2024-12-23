package models

import (
	"time"
)

type Users struct {
	ID       int    `gorm:"type:int;primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     string `gorm:"type:varchar(255)" json:"role"`

	//time stamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
