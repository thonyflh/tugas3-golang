package model

type Product struct {
	BaseStruct
	Name   string `json:"name" gorm:"type:varchar(255)"`
	Detail string `json:"detail" gorm:"type:varchar(255)"`
	Type   string `json:"type" gorm:"type:varchar(255)"`
	Stock  int    `json:"stock" gorm:"type:smallint"`
}
