package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"unique; primaryKey; autoIncrement"`
	Login    string `json:"login" gorm:"type: varchar(32); unique; not null"`
	Username string `json:"username" gorm:"type: varchar(32)"`
	Password string `json:"password" gorm:"type: varchar(32)"`
}

type Category struct {
	gorm.Model
	Id           uint   `json:"id" gorm:"unique; primaryKey; autoIncrement"`
	CategoryName string `json:"category_name" gorm:"type:varchar(64); unique; not null"`
}

type Product struct {
	gorm.Model
	Id           uint            `json:"id" gorm:"unique; primaryKey; autoIncrement"`
	ProductName  string          `json:"product_name" gorm:"type:varchar(64); unique; not null"`
	Price        decimal.Decimal `json:"price" gorm:"type: decimal(10, 2); not null"`
	Description  string `json:"description" gorm:"type: text"`
	CategoryName Category `json:"category_name" gorm:"foreignKey:CategoryName"`
}
