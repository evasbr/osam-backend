package model

import "gorm.io/datatypes"

type User struct {
	BaseModel
	Email           string          `gorm:"unique;not null"`
	Password        string          `gorm:"not null" json:"-"`
	Name            string          `gorm:"not null"`
	BirthDate       *datatypes.Date `gorm:"type:date"`
	TelephoneNumber *string
}
