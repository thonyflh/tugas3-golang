package model

import "time"

type User struct {
	BaseStruct
	Name     string     `json:"name" gorm:"type:varchar(255)"`
	Phone    string     `json:"phone" gorm:"type:varchar(20)"`
	Address  string     `json:"address" gorm:"type:varchar(255)"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Email    string     `json:"email" gorm:"type:varchar(255)"`
}
