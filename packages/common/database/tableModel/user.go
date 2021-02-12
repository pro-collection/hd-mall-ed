package tableModel

import (
	"gorm.io/gorm"
)

type BaseUser struct {
	Name             string `json:"name" gorm:"not null"`
	Integral         uint   `json:"integral" gorm:"comment:积分;index"`
	DefaultAddressId uint   `json:"default_address_id" gorm:"comment:默认地址"`
	Avatar           string `json:"avatar" gorm:"comment:头像"`
	Email            string `json:"email" gorm:"comment:邮箱"`
	Role             string `json:"role" gorm:"comment:角色"`
}

type User struct {
	gorm.Model
	Password string `json:"password" gorm:"not null"`
	BaseUser
}
