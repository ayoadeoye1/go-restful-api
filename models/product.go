package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Products struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Name        string      `gorm:"type:varchar(255);not null" json:"name"`
	Description string      `gorm:"type:text;not null" json:"description"`
	Price       int         `gorm:"type:int;not null" json:"price"`
	Currency    string      `gorm:"type:varchar(255);not null" json:"currency"`
	Category    string      `gorm:"type:varchar(255)" json:"category"`
	Brand       *string     `gorm:"type:varchar(255)" json:"brand"`
	Stock       uint        `gorm:"type:int" json:"stock"`
	Images      StringArray `gorm:"type:json" json:"images"`

	// Time Stamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Products) TabelName() string {
	return "products"
}
