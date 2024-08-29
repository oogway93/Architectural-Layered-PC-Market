package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Login    string `json:"login" gorm:"type:varchar(32);unique;not null"`
	Username string `json:"username" gorm:"type:varchar(32)"`
	Password string `json:"password" gorm:"type:varchar(32)"`
}

type Category struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CategoryName string    `json:"category_name" gorm:"type:varchar(64);unique;not null"`
	Products     []Product `gorm:"foreignKey:CategoryID"`
	// Products     []Product `json:"products,omitempty" gorm:"polymorphic:Category"`
}

// models/product.go
type Product struct {
	gorm.Model
	ID          uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UUID        uuid.UUID       `json:"uuid" gorm:"type:uuid;default:gen_random_uuid();index"`
	ProductName string          `json:"product_name" gorm:"type:varchar(64);unique;not null"`
	Price       decimal.Decimal `json:"price" gorm:"type:decimal(10, 2);not null"`
	Description string          `json:"description" gorm:"type:text"`
	CategoryID  uint            `json:"categoryId" gorm:"foreignKey:CategoryID"`
}
