package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Login    string `json:"login" gorm:"type:varchar(32);unique;not null"`
	Username string `json:"username" gorm:"type:varchar(32)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}

type Category struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CategoryName string    `json:"category_name" gorm:"type:varchar(64);unique;not null"`
	Products     []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	ID          uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UUID        uuid.UUID       `json:"uuid" gorm:"type:uuid;default:gen_random_uuid();index"`
	ProductName string          `json:"product_name" gorm:"type:varchar(64);unique;not null"`
	Price       decimal.Decimal `json:"price" gorm:"type:decimal(10, 2);not null"`
	Description string          `json:"description" gorm:"type:text"`
	CategoryID  uint            `json:"categoryId" gorm:"index"`
	Category    Category        `json:"category" gorm:"foreignKey:CategoryID"`
}

type Order struct {
	gorm.Model
	ID         uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserID     uint            `json:"userId" gorm:"index"`
	Status     string          `gorm:"type:varchar(20);not null;default:'pending'" validate:"oneof=pending in_process shipped delivered picked_up completed cancelled"`
	Total      decimal.Decimal `json:"total" gorm:"type:decimal(10, 2)"`
	DeliveryID uint            `json:"deliveryId" gorm:"index"`

	OrderItems []*OrderItem `gorm:"foreignKey:OrderID"`
	User       User         `gorm:"foreignKey:UserID"`
	Delivery   Delivery     `gorm:"foreignKey:DeliveryID"`
}

type OrderItem struct {
	gorm.Model
	ID        uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	OrderID   uint            `json:"orderId" gorm:"index"`
	ProductID uint            `json:"productId" gorm:"index"`
	Quantity  int             `json:"quantity" gorm:"not null"`
	UnitPrice decimal.Decimal `json:"unit_price" gorm:"type:decimal(10, 2)"`

	Order   *Order  `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

type Delivery struct {
	gorm.Model
	ID            uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	FullName      string          `json:"full_name" gorm:"type varchar(64)"`
	Postcode      string          `json:"postcode" gorm:"type varchar(64)"`
	Country       string          `json:"country" gorm:"type:varchar(64)"`
	City          string          `json:"city" gorm:"type:varchar(64)"`
	DeliveryPrice decimal.Decimal `json:"delivery_price" gorm:"type:decimal(10, 2)"`
}
