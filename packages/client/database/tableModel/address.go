package tableModel

import "gorm.io/gorm"

type BaseAddress struct {
	ID            uint   `gorm:"primarykey"`
	UserId        uint   `json:"user_id" gorm:"not null;index;comment:'关联的userId'"`
	Province      string `json:"province" gorm:"not null;comment:'省市区'"`
	AddressDetail string `json:"address_detail" gorm:"not null;comment:'地址详情'"`
	AddresseeName string `json:"addressee_name" gorm:"not null;comment:'收件人姓名'"`
	Mobile        string `json:"mobile" gorm:"not null;comment:'收件人邮箱'"`
}

type Address struct {
	gorm.Model
	BaseAddress
}
