package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `json:"id" gorm:"column id;primaryKey; not null; autoIncrement"`
	Username string `json:"username" gorm:"column:username; type: varchar(32); unique; not null; check: "`
	Password string `json:"password" gorm:"column: password; type: varchar(32); not null"`
}