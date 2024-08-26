package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"unique; primaryKey; autoIncrement"`
	Login    string `json:"login" gorm:"type: varchar(32); unique; not null"`
	Username string `json:"username" gorm:"type: varchar(32)"`
	Password string `json:"password" gorm:"type: varchar(32)"`
}
