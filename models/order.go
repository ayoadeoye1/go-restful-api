package models

import "time"

type Order struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID      uint   `gorm:"not null"`
	TotalAmount int    `gorm:"not null"`
	Address     string `gorm:"type:varchar(255);not null'"`
	Status      string `gorm:"type:varchar(50);not null;default:'Pending'"`
	// 'Pending', 'Processing', 'Completed', 'Cancelled'

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User Users `gorm:"foreignKey:UserID"`
}

type OrderItem struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	OrderID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  uint `gorm:"not null"`
	Price     int  `gorm:"not null"`
	Subtotal  int  `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Order   Order    `gorm:"foreignKey:OrderID"`
	Product Products `gorm:"foreignKey:ProductID"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderItem) TableName() string {
	return "orderitems"
}
