package models

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	MailID       string `json:"mailID" gorm:"not null"`
	PasswordHash string `json:"passwordHash" gorm:"not null"`
	RandomHash   string `json:"randomHash" gorm:"not null"`

	Name        string  `json:"name" gorm:"not null"`
	StudID      string  `json:"studID" gorm:"not null"`
	Designation string  `json:"designation"`
	Batch       int     `json:"batch"`
	Address     string  `json:"address"`
	CountryID   Country `json:"countryID" gorm:"foreignKey:ID"`
	MobileNo    string  `json:"mobileNo"`
}
