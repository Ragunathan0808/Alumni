package models

type Country struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt uint64 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt uint64 `json:"updatedAt" gorm:"autoUpdateTime"`

	MailID       string `json:"mailID" gorm:"unique;not null"`
	PasswordHash string `json:"passwordHash" gorm:"not null"`
	Verified     bool   `json:"verified" gorm:"default:false"`
	Active       bool   `json:"active" gorm:"default:false"`
	Type         string `json:"type" gorm:"default:user"`

	Name        string  `json:"name" gorm:"not null"`
	StudID      string  `json:"studID" gorm:"unique;not null"`
	Designation string  `json:"designation" gorm:"index"`
	Batch       uint    `json:"batch" gorm:"index"`
	Address     string  `json:"address"`
	CountryID   Country `json:"countryID" gorm:"foreignKey:ID"`
	MobileNo    string  `json:"mobileNo"`
}
