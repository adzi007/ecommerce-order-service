package model

import (
	"time"
)

type Order struct {
	ID            uint          `gorm:"primaryKey"`
	UserId        string        `gorm:"not null"`
	TotalPrice    float64       `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string        `gorm:"type:varchar(50);default:null"`
	PaymentFee    float64       `gorm:"type:decimal(10,2);default:0"`
	Status        string        `gorm:"type:varchar(50);default:'Pending'"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
