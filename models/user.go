package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);default:null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Phone     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Status    string    `gorm:"type:enum('ACTIVE','INACTIVE','SUSPENDED');default:ACTIVE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}
